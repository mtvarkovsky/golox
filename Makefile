EXEC_NAME=golox
BUILD_DIR=./build
BUILD_SRC=./cmd/$(EXEC_NAME).go
BUILD_OUT=$(BUILD_DIR)/$(EXEC_NAME)

args = $(foreach a,$($(subst -,_,$1)_args),$(if $(value $a),$a="$($a)"))

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## builds the executable and places it to ./build/
	go build -o ${BUILD_OUT} ${BUILD_SRC}

.PHONY: generate
generate: ## runs go generate ./...
	go generate ./...

.PHONY: test
test: ## runs tests
	go test ./...

.PHONY: golox
golox: build ## runs the interpreter executable from ./build/
	${BUILD_OUT}
