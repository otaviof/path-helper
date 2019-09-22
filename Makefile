# application name
APP = path-helper
# application version
VERSION ?= $(shell cat ./version)
# build directory
BUILD_DIR ?= build
# end-to-end directory
E2E_DIR ?= test/e2e
# build flags
BUILD_FLAGS ?= -v -a -ldflags=-s -mod=vendor
# gopath copy from environment
GOPATH ?= ${GOPATH}

.PHONY: default bootstrap build clean test

default: build

bootstrap:
	GO111MODULE=on go mod vendor

build: clean
	GO111MODULE=on CGO_ENABLED=0 go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP) cmd/$(APP)/*

install: build
	GO111MODULE=on CGO_ENABLED=0 go install $(BUILD_FLAGS) cmd/$(APP)/*

clean:
	rm -rf $(BUILD_DIR) > /dev/null

clean-vendor:
	rm -rf ./vendor > /dev/null

test: test-unit test-e2e

test-unit:
	go test -failfast -race -coverprofile=coverage.txt -covermode=atomic -cover -v pkg/$(APP)/*

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
	git push origin $(VERSION)
	goreleaser --rm-dist
