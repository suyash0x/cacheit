.PHONY:	build test

build:
	@go build -o bin/cacheit cacheit.go

test:
	@go test -race