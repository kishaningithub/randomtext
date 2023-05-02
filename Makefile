HOSTOS=$(shell uname | tr "[:upper:]" "[:lower:]")
VERSION=$(shell git tag --sort=-version:refname | head -n 1)

TOOLS_DIR=tools
TOOLS_VERSION_NANCY=v1.0.21

NANCY=$(TOOLS_DIR)/nancy-$(TOOLS_VERSION_NANCY)

NANCY_EXISTS:=$(shell test -e "$(NANCY)" && echo -n "binary exists")

setup: ## Setup necessary binary dependencies
	mkdir -p $(TOOLS_DIR)
ifndef NANCY_EXISTS
	wget -O $(TOOLS_DIR)/nancy-$(TOOLS_VERSION_NANCY) https://github.com/sonatype-nexus-community/nancy/releases/download/$(TOOLS_VERSION_NANCY)/nancy-$(TOOLS_VERSION_NANCY)-$(HOSTOS)-amd64
	chmod +x $(TOOLS_DIR)/nancy-$(TOOLS_VERSION_NANCY)
endif

fmt: ## Run the code formatter
	gofmt -l -s -w .

download-deps:
	go mod download

tidy-deps:
	go mod tidy

compile:
	go build

unit-test:
	go test -v ./...

dependency-check: ## Ensure dependencies have no vulnerabilities
	go list -json -m all | $(NANCY) sleuth --skip-update-check

build: setup download-deps tidy-deps fmt unit-test compile

update-deps:
	go get -u -t ./...
	go mod tidy
