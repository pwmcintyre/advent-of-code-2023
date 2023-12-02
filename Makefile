SHELL := /bin/bash

# This version-strategy uses git tags to set the version string
VERSION ?= $(shell git describe --tags --always --dirty)

SRC := $(shell find . -type f -name '*.go')

build: main.wasm
main.wasm: $(SRC)
	GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go

run: main.wasm
	wasmtime main.wasm

days/1/main.wasm: days/1/main.go
	GOOS=wasip1 GOARCH=wasm go build -o days/1/main.wasm days/1/main.go
