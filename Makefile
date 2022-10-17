APP = path-helper
OUTPUT_DIR ?= _output
VERSION ?= $(shell cat ./version)

BIN ?= $(OUTPUT_DIR)/$(APP)
CMD ?= cmd/$(APP)/*
PKG ?= pkg/$(APP)/*

E2E_DIR ?= test/e2e
GOFLAGS_TEST ?= -failfast -race -coverprofile=coverage.txt -covermode=atomic -cover -v
GOFLAGS ?= -v -a -ldflags=-s -mod=vendor

ARGS ?=

.EXPORT_ALL_VARIABLES:

default: build

.PHONY: $(BIN)
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
	rm -rf $(OUTPUT_DIR) > /dev/null

test: test-unit test-e2e

.PHONY: test-unit
test-unit:
	go test $(GOFLAGS_TEST) pkg/$(APP)/*

.PHONY: test-e2e
test-e2e:
	./test/e2e/bats/core/bin/bats --recursive $(E2E_DIR)/*.bats

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
