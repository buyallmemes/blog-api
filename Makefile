pre-build:
	go mod tidy
test:
	go test ./...
build: pre-build test
	go build
