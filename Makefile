##
# Makefile configuration
# See: https://www.gnu.org/software/make/manual/make.html
.ONESHELL:

SHELL = /bin/bash
SHELLFLAGS = -ce

.DEFAULT_GOAL = help

.PHONY: start
start: ## Run the server locally.
	air

.PHONY: build
build: ## Build the application from source.
	go build -ldflags="-w -s" -o tmp/main ./cmd

.PHONY: zbuild
zbuild: build ## Compresses the build output.
	upx --best tmp/main

.PHONY: unit
unit: ## Run the unit test suite.
	go test ./...

.PHONY: lint
lint: ## Lint this projects source files.
	revive -config revive.toml ./...

.PHONY: help
help:  ## Print this help.
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sed 's/Makefile://' | \
		sort -d | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Silence output by default, use `VERBOSE=1 make <command>` to enable.
ifndef VERBOSE
.SILENT:
endif
