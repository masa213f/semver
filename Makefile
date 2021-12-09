PROJECT_DIR := $(CURDIR)
BIN_DIR := $(PROJECT_DIR)/bin
DIST_DIR := $(PROJECT_DIR)/dist

GOIMPORTS := $(BIN_DIR)/goimports
STATICCHECK := $(BIN_DIR)/staticcheck
GORELEASER := $(BIN_DIR)/goreleaser

.PHONY: all
all: help

##@ Basic
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: setup
setup: ## Setup necessary tools.
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install golang.org/x/tools/cmd/goimports@latest
	GOBIN=$(BIN_DIR) go install honnef.co/go/tools/cmd/staticcheck@latest
	GOBIN=$(BIN_DIR) go install github.com/goreleaser/goreleaser@latest

.PHONY: clean
clean: ## Clean files
	-rm $(BIN_DIR)/* $(DIST_DIR)/*
	-rmdir $(BIN_DIR) $(DIST_DIR)

##@ Build

.PHONY: build
build: ## Build a binary.
	$(GORELEASER) build --snapshot --rm-dist --single-target

.PHONY: release-build
release-build:
	$(GORELEASER) build --snapshot --rm-dist

.PHONY: format
format: ## Format go files.
	$(GOIMPORTS) -w $$(find . -name '*.pb.go' -prune -o -name '*.go' -print)

##@ Test

.PHONY: lint
lint:
	test -z "$$($(GOIMPORTS) -l $$(find . -name '*.pb.go' -prune -o -name '*.go' -print) | tee /dev/stderr)"
	$(STATICCHECK) ./...
	go vet ./...

.PHONY: test
test: ## Run unit tests.
	go test -v ./...
