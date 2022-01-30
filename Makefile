# Adds $GOPATH/bit to $PATH
export PATH := $(PATH):$(shell go env GOPATH)/bin

help: ## Display this help message
	@echo "Usage:"
	@grep -E "^[a-zA-Z_-]+:.*? ## .+$$" $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}'

lint: ## Run golangci-lint
	golangci-lint run

lint-install: ## Install golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0

.PHONY: help lint lint-install
