# An inchoate Makefile for building the Go binaries.
#
# If you're tweaking proto files, you must run "make proto" manually otherwise
# your changes to the .proto file will not get generated.
#

# The name of the package containing Smashwords code, as Go understands it.
PACKAGE=miceplans.net

# The path of the root of the package.
PACKAGE_SRC=.

# The version of Go expected for deployment.
GO_VERSION=go version go1.13.1 linux/amd64

export GOBIN=${PWD}/bin
# Force go to use files in vendor dir

# builds everything
default:
	go install

# Show the Go version expected for deployment.
goversion: ; @echo $(GO_VERSION)

# clean removes all compiled binaries and compiled dependencies.
clean:
	go clean
