APP = path-helper
OUTPUT_DIR ?= _output
VERSION ?= $(shell cat ./version)

BIN ?= $(OUTPUT_DIR)/$(APP)
CMD ?= cmd/$(APP)/*
PKG ?= pkg/$(APP)/*

GOPATH ?= ${GOPATH}

E2E_DIR ?= test/e2e
GO_FLAGS ?= -v -a -ldflags=-s -mod=vendor

ARGS ?=

default: build

.PHONY: vendor
vendor:
	go mod vendor

$(BIN):
	go build $(GO_FLAGS) -o $(BIN) $(CMD)
build: $(BIN)

.PHONY: run
run:
	go run $(GO_FLAGS) $(CMD) $(ARGS)

install: build
	install -m 0755 $(BIN) $(GOPATH)/bin/

.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR) > /dev/null

test: test-unit test-e2e

.PHONY: test-unit
test-unit:
	go test -failfast -race -coverprofile=coverage.txt -covermode=atomic -cover -v pkg/$(APP)/*

.PHONY: test-e2e
test-e2e:
	bats --recursive $(E2E_DIR)

hack-install-bats:
	hack/install-bats.sh

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
