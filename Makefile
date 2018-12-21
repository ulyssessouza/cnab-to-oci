.DEFAULT_GOAL := all

SHELL:=/bin/bash

all: build test

all-ci: lint all

get-tools:
	go get github.com/alecthomas/gometalinter
	gometalinter --install

build:
	go build -o bin/cnab-to-oci ./cmd/cnab-to-oci

install:
	pushd cmd/cnab-to-oci && go install && popd

clean:
	rm -rf bin

test:
	go test -failfast ./...

format:
	go fmt ./...
	goimports -w .

lint: get-tools
	gometalinter --config=gometalinter.json ./...

.PHONY: all, get-tools, build, clean, test, lint
