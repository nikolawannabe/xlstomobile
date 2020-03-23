package main

import (
	"log"
	"net/http"

	xlsxparser "github.com/nikolawannabe/xlsxparser"
)

func getInventoryList(w http.ResponseWriter, r *http.Request) {
	parser, err := xlsxparser.NewParser("/root/animekat/testfile.xlsx")
	if err != nil {
		panic(err)
	}
	products, _ := parser.ParseFile()

	output, err := getHtmlString(products)
	if err != nil {
		log.Printf("unable to run template: %v", err)
	}
	w.Write([]byte(output))
}

func main() {
	http.HandleFunc("/getInventory", getInventoryList)
	log.Fatal(http.ListenAndServe(":80", nil))
}
