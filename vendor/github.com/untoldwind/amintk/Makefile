PKGS = ./fixtures/... ./glib/... ./gdk/... ./gtk/...

all: export GOPATH=${PWD}/../../../..
all:
	@echo "--> Build"
	@go build -i ${PKGS}

bin.examples: export GOPATH=${PWD}/../../../..
bin.examples:
	@echo "--> Build examples"
	@go build -i -o bin/simple ./examples/simple

format: export GOPATH=${PWD}/../../../..
format:
	@echo "--> Running go fmt"
	@go fmt ${PKGS}

test: export GOPATH=${PWD}/../../../..
test:
	@echo "--> Running tests"
	@go test -v ${PKGS}

glide.install:
	@echo "--> glide install"
	@go get github.com/Masterminds/glide
	@go build -v -o bin/glide github.com/Masterminds/glide
	@bin/glide install -v
