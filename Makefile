BINARY_NAME=glvd

.PHONY: all fmt build clean build-linux build-linux-amd64 build-linux-arm64

all: fmt build

fmt:
	go fmt ./...

build:
	go build -o $(BINARY_NAME) .

build-linux: build-linux-amd64 build-linux-arm64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 .

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o $(BINARY_NAME)-linux-arm64 .

clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME)-linux-amd64 $(BINARY_NAME)-linux-arm64
