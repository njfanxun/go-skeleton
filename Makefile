export SHELL:=/bin/bash
export SHELL_OPTS:=$(if $(SHELL_OPTS),$(SHELL_OPTS):)pipefail:errexit
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
.SUFFIXES:
APP_NAME := gosk
override LDFLAGS = ""
CGO_ENABLED ?= 1
WASM_ENABLED ?= 1

GO := CGO_ENABLED=$(CGO_ENABLED) GO111MODULE=on go
GOOS ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)


build:
	$(GO) build -o ./bin/$(APP_NAME)  main.go

.PHONY: build-bin
build-bin:
	@$(MAKE) GOOS=$(GOOS) GOARCH=$(GOARCH) build WASM_ENABLED=0 CGO_ENABLED=0