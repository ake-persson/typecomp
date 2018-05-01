all:	build

format:
	gofmt -w .

test:
	golint -set_exit_status .
	go vet ./...
	errcheck ./...
	go test ./... -v -covermode=atomic

build: format test
	go build

.PHONY: format test build
