# Basic Makefile
build:
	echo $$GOPATH
	go get -d
	go build