build: lint
	go build .

lint:
	golangci-lint run

test:
	go test -v ./pkg/...
