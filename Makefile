PROJECT=paymill-client-go

BUILD_PATH := $(shell pwd)/.gobuild

C0_PATH := "$(BUILD_PATH)/src/github.com/catalyst-zero"

BIN := $(PROJECT)

.PHONY=clean run-test get-deps update-deps

GOPATH := $(BUILD_PATH)

SOURCE=$(shell find . -name '*.go')

all: get-deps $(BIN)

clean:
	rm -rf $(BUILD_PATH) $(BIN)

get-deps: .gobuild

.gobuild:
	mkdir -p $(C0_PATH)
	cd "$(C0_PATH)" && ln -s ../../../.. $(PROJECT)

	#
	# Fetch private packages first (so `go get` skips them later)

	#
	# Fetch public dependencies via `go get`
	GOPATH=$(GOPATH) go get -d -v github.com/catalyst-zero/$(PROJECT)

	#
	# Fetch deployment dependencies
	#GOPATH=$(GOPATH) go get github.com/kr/godep

	#
	# Build test packages (we only want those two, so we use `-d` in go get)
	#GOPATH=$(GOPATH) go get -v github.com/onsi/gomega
	#GOPATH=$(GOPATH) go get -v github.com/onsi/ginkgo

$(BIN): $(SOURCE)
	GOPATH=$(GOPATH) go build -o $(BIN)

run-tests:
	GOPATH=$(GOPATH) go test ./...

fmt:
	gofmt -l -w .
