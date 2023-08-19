APP = path-helper
OUTPUT_DIR ?= bin
VERSION ?= $(shell cat ./version)

BIN ?= $(OUTPUT_DIR)/$(APP)
CMD ?= ./cmd/$(APP)/...
PKG ?= ./pkg/$(APP)/...

GOFLAGS ?= -v -a
CGO_LDFLAGS ?= -s -w

GOFLAGS_TEST ?= -failfast -race -coverprofile=coverage.txt -covermode=atomic -cover -v

BATS_CORE ?= test/e2e/bats/core/bin/bats
E2E_DIR ?= test/e2e
E2E_TEST_GLOB ?= *.bats
E2E_TESTS = $(E2E_DIR)/$(E2E_TEST_GLOB)

ARGS ?=

.EXPORT_ALL_VARIABLES:

default: build

$(BIN):
	go build -o $(BIN) $(CMD)

build: $(BIN)

.PHONY: run
run:
	go run $(CMD) $(ARGS)

install: build
	install -m 0755 $(BIN) $(GOPATH)/bin/

.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR) >/dev/null

test: test-unit test-e2e

.PHONY: test-unit
test-unit:
	go test $(GOFLAGS_TEST) $(CMD) $(PKG)

.PHONY: test-e2e
test-e2e: $(BIN)
	$(BATS_CORE) --trace --verbose-run --recursive $(E2E_TESTS)

codecov:
	mkdir .ci || true
	curl -s -o .ci/codecov.sh https://codecov.io/bash
	bash .ci/codecov.sh -t $(CODECOV_TOKEN)

snapshot:
	goreleaser --rm-dist --snapshot

snapshot-local:
	goreleaser --rm-dist --snapshot --skip-publish --debug

snapshot-install: snapshot-local
	install -m 755 build/dist/darwin_darwin_amd64/$(APP) "$(GOPATH)/bin/$(APP)"

release:
	git tag $(VERSION)
	git push --tags origin $(VERSION)
	goreleaser --rm-dist
