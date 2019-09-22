# application name
APP = path-helper
# application version
VERSION ?= $(shell cat ./version)
# build directory
BUILD_DIR ?= build

.PHONY: default bootstrap build clean test

default: build

bootstrap:
	GO111MODULE=on go mod vendor

build: clean
	go build -v -o $(BUILD_DIR)/$(APP) .

clean:
	rm -rf $(BUILD_DIR) > /dev/null

clean-vendor:
	rm -rf ./vendor > /dev/null

test:
	go test -failfast -race -coverprofile=coverage.txt -covermode=atomic -cover -v .

codecov:
	mkdir .ci || true
	curl -s -o .ci/codecov.sh https://codecov.io/bash
	bash .ci/codecov.sh -t $(CODECOV_TOKEN)

snapshot:
	goreleaser --rm-dist --snapshot

release:
	git tag $(VERSION)
	git push origin $(VERSION)
	goreleaser --rm-dist
