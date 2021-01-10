.PHONY: build deploy test

build:
	CGO_ENABLED=0  go build -o ./dist/sur -a -ldflags '-s' -installsuffix cgo .

test:
	mkdir -p static/packages
	cp test/fixtures/eopkg-index.xml static/packages
	go test ./...

