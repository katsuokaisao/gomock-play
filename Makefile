GOVERSION=$(shell go version)
THIS_GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
THIS_GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
GOOS?=$(THIS_GOOS)
GOARCH?=$(THIS_GOARCH)

DIR_BUILD=build
DIR_CURRENT=$(shell pwd)

SRC:=$(shell find . -type f -name '*.go')

.PHONY: \
	build \
	clean \
	test

$(DIR_BUILD)/$(GOOS)_$(GOARCH)/gomock: $(SRC) go.mod go.sum
	@mkdir -p $(DIR_BUILD)/$(GOOS)_$(GOARCH)
	@echo "Building $@..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o $@ ./*.go

build: $(DIR_BUILD)/$(GOOS)_$(GOARCH)/gomock

clean:
	@rm -rf $(DIR_BUILD)
	@rm -rf coverage.out

test:
	@go test -p 1 -v ./...

coverage:
	@go test -p 1 -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out