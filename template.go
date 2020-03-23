package main

import (
	"bytes"
	"text/template"

	xlsxparser "github.com/nikolawannabe/xlsxparser"
)

func getHtmlString(productTypes xlsxparser.ProductTypes) (string, error) {
	const bootstrapCdn = `<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">`
	const tpl = `
	{{define "foo"}}
	<html>
		<head>
		{{ .Bootstrap }}
		</head>
		<body>
		{{ range $index, $element := .Types}}
			<h1>{{$index}}</h1>
			{{ with $element }}
				<div class="container">

				{{ range $rowNum, $row :=  .ProductRows }}
				{{ if mod $rowNum 2 }}
				<div class="row">
				{{ end }}
						<div class="border rounded m-2 col-md">
						{{ with $row }}
							  <span class="font-weight-bold">{{ .ProductName }}</span></br>
								<p>Category: {{ .Category}}</p>
								<p class="text-muted" style="height: 60px;">{{ .StockLevel }} left</p>
								<span class="text-center" style="position: absolute; bottom: 0; height 20px;">
									<span class="badge badge-pill badge-warning my-2">
										${{ .Price }}
									</span>
								</span>
						{{ end }}
						</div>
				{{ if mod1 $rowNum 2  }}
				</div>
				{{ end }}
				{{ end }}
				</div>
				</div>
			{{ end }}
		{{ end}}
		</body>
	</html>
	{{end}}`
	funcMap := template.FuncMap{
		"mod":  func(i, j int) bool { return i%j == 0 },
		"mod1": func(i, j int) bool { return (i-1)%j == 0 },
	}
	t, err := template.New("foo").Funcs(funcMap).Parse(tpl)
	if err != nil {
		return "", err
	}

	data := struct {
		Bootstrap string
		Types     xlsxparser.ProductTypes
	}{
		Bootstrap: bootstrapCdn,
		Types:     productTypes,
	}

	var buf bytes.Buffer
	err = t.ExecuteTemplate(&buf, "foo", data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
