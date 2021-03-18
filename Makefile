PROJ = "epn"
TIME = $$(date +%Y-%m-%d_%H:%M)
COMMIT = $$(git rev-parse --short HEAD)
VERSION = $$(git describe --exact-match --tags 2> /dev/null || git rev-parse --short HEAD)
BIN = ./bin/${PROJ}-${VERSION}.bin
ARGS ?= ""
FAKEDATA_N ?= 1
OSM_ZONE ?= asia/taiwan

.PHONY: FORCE

##@ Show

version:  ## Show version
	echo ${VERSION}

.PHONY: version

##@ Build

force-build: FORCE
	GOOS=linux go build -o ./bin/${PROJ}-${VERSION}.bin  -ldflags "-s -w -X main.date=${TIME} -X main.commit=${COMMIT} -X main.version=${VERSION}" ./*.go
force-build-win: FORCE
	GOOS=windows go build -o ./bin/${PROJ}.exe  -ldflags "-s -w -X main.date=${TIME} -X main.commit=${COMMIT} -X main.version=${VERSION}" ./*.go

build: force-build ## Build binary
build-win: force-build-win ## Build binary for windows

.PHONY: build, force-build-win

##@ Testing

bench-test:  ## Run go test
	go test -v -cover -bench=. -benchtime=2x ./... | tee | richgo testfilter

install-richgo:  ## Install richgo
	go get -u github.com/kyoh86/richgo

.PHONY: test install-richgo bench-test
##@ Lint

install-lint:  ## Install golangci-lint binary to ./bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0

lint-testing:  ## Run lint
	./bin/golangci-lint run ./...

.PHONY: install-lint lint-testing

##@ Help

.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
