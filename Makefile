build: lint
	go build .

test: lint
	go clean -testcache
	go test -v ./...

lint:
	golangci-lint run
