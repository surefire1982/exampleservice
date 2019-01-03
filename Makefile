# Basic Makefile
.PHONY: all
all: build
FORCE: ;

.PHONY: clean build

clean:
	rm -rf bin/*

dependencies:
	go get -d ./...

build: dependencies build-service
	
build-service:
	go build -o ./bin/exampleservice api/main.go

test:
	go test -v -tags testing ./...