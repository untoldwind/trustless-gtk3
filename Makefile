PKGS= ./ui/... .

VERSION ?= $(shell date -u +%Y%m%d.%H%M%S)VERSION ?= $(shell date -u +%Y%m%d.%H%M%S)

#all: export GOPATH=${PWD}/../../../..
all: format
	@mkdir -p bin
	@echo "--> Running go build ${VERSION}"
	@go build -v -i -o bin/trustless-gkt3 github.com/untoldwind/trustless-gtk3

#format: export GOPATH=${PWD}/../../../..
format:
	@echo "--> Running go fmt"
	@go fmt ${PKGS}

glide.install:
	@echo "--> glide install"
	@go get github.com/Masterminds/glide
	@go build -v -o bin/glide github.com/Masterminds/glide
	@bin/glide install -v
