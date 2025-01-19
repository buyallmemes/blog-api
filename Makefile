NAME=blog-api
BUILD_FOLDER=build

made deps:
	go mod tidy
	go mod vendor
pre-build:
	go mod tidy
go-test:
	go test -race ./...
clean:
	rm -rf build/
go-build:
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ${BUILD_FOLDER}/
	zip ${BUILD_FOLDER}/${NAME}.zip ${BUILD_FOLDER}/${NAME}

sam-build:
	sam build

sam-deploy:
	sam deploy

build: clean go-test go-build sam-build
