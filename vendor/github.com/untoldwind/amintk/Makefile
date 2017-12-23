PKGS = ./cairo/... ./fixtures/... ./glib/... ./gdk/... ./gtk/...

all: export GOPATH=${PWD}/../../../..
all: format
	@echo "--> Build"
	@go install ${PKGS}

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
	@go test -i -v ${PKGS}
	@go test -v ${PKGS}

bin/dep:
	@echo "-> dep install"
	@go get github.com/golang/dep/cmd/dep
	@go build -v -o bin/dep github.com/golang/dep/cmd/dep

dep.ensure: bin/dep
	@bin/dep ensure
	@bin/dep prune
	@find vendor -name "*_test.go" -exec rm -f {} \;
	@find vendor -type f ! -name "*.go" -exec rm -f {} \;
