.PHONY: update build clean binary tests

clean:
	go mod tidy
	go mod vendor -e

update:
	-GOFLAGS="" go get all

build:
	GOOS=js GOARCH=wasm go build ./...

tests:
	GOOS=js GOARCH=wasm go test -v ./...

test: tests
