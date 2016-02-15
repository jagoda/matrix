default: build

build: prepare
	godep go build

lint:
	go fmt ./...
	go vet ./...

prepare:
	go get github.com/tools/godep

test: prepare lint
	godep go test -cover ./...
