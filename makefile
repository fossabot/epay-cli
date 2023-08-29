VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse HEAD)
DATE := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	@go build -o bin/$(BINARY_NAME) -v -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"
