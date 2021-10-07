build: lint
	go build .

test: lint
	go test -v ./...

lint:
	golangci-lint run
