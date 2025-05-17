NAME=blog-api

deps:
	@go mod tidy
	@go mod vendor
pre-build:
	@go mod tidy
go-test:
	@go test -race ./...
	@go vet ./...
clean:
	@rm -rf .aws-sam/

build-BlogAPI:
	@GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o $(ARTIFACTS_DIR)/bootstrap
	@cp -R resources $(ARTIFACTS_DIR)

build: clean go-test
	@sam build

update:
	@go get -u ./...
	@go mod tidy
	@go mod vendor

run: build
	@sam local start-api