.PHONY: build test dev

VERSION=1.2.0

build:
	CGO_ENABLED=0  go build -o ./dist/sur -a -ldflags '-s' -installsuffix cgo .

test:
	mkdir -p static/packages
	cp test/fixtures/eopkg-index.xml static/packages

	SUR_PACKAGES_DIR=static/packages \
	go test ./...

dev:
	SUR_PACKAGES_DIR=static/packages \
	go run .
