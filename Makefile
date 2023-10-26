SHELL=/bin/bash

GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean

BINARY_NAME=cfx
CMD_PATH=./main.go
DIST=dist
DIST_LINUX=$(DIST)/linux

MODULE=$(shell go list -m)
DATE=$(shell date +%D_%H:%M)

build-linux:
	mkdir -p $(DIST_LINUX)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -ldflags '-X $(MODULE)/main.BuildDate=$(DATE)' -o ./$(DIST_LINUX)/$(BINARY_NAME) -v $(CMD_PATH)

build-linux-obfuscated:
	mkdir -p $(DIST_LINUX)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 garble build -ldflags '-X $(MODULE)/main.BuildDate=$(DATE)' -o ./$(DIST_LINUX)/$(BINARY_NAME)-obfuscated -v $(CMD_PATH)

build: build-linux

build-obfuscated: build-linux-obfuscated

clean:
	rm -rf $(DIST)
