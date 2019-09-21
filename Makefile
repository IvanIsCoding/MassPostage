VERSION := 0.1
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

build:
	go build -o bin/masspost $(GOFILES)

clean:
	go clean -i .
	rm -rf bin