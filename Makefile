PACKAGES := $(shell go list ./pkg...)

pre-build:
	go mod tidy
test:
	go test -v $(PACKAGES)
build: pre-build test
	go build
