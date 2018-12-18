all: binary test

binary:
	go build -o bin/cnab-to-oci github.com/docker/cnab-to-oci/cmd/cnab-to-oci

.PHONY: test
test:
	go test ./...