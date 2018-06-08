# Define parameters
BINARY=automata
SHELL := /bin/bash
GOPACKAGES = $(shell go list ./... | grep -v vendor)
ROOTDIR = $(pwd)

.PHONY: build install test linux

default: build

build: automata.go
	go build -v -o ./build/${BINARY} automata.go

install:
	go install  ./...

test:
	go test -race -cover ${GOPACKAGES}

clean:
	rm -rf build

linux: automata.go
	GOOS=linux GOARCH=amd64 go build -o ./build/linux/${BINARY} automata.go
