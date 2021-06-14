.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

deps: ## Get all the required dependencies
	go mod tidy

build-linux: deps ## Build CLI for Linux
	GOOS=linux go build -o turing src/

build-win: deps ## Build CLI for Linux
	GOOS=windows go build -o turing src/

buildall: build-linux build-win ## Build all supported versions of the CLI