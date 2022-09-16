.DEFAULT_GOAL := test

build: lint
	go build .

test: lint install-gotestfmt
	go clean -testcache
	go test -json -v ./...  2>&1 | gotestfmt

lint:
	golangci-lint run

install-gotestfmt:
ifeq (, $(shell command -v gotestfmt 2> /dev/null))
	go install github.com/haveyoudebuggedit/gotestfmt/v2/cmd/gotestfmt@latest
endif
