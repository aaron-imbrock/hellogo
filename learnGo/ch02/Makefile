.DEFAULT_GOAL := build

SRC := $(wildcard *.go)
EXECUTABLES := $(SRC:.go=)

fmt:
		go fmt ./...

lint: fmt
		golint ./...

vet: fmt
		go vet ./...

build: vet $(EXECUTABLES)
.PHONY:build

%: %.go
	go build -o $@ $<
.PHONY $(EXECUTABLES)