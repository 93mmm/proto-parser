BUILD_NAME=protosym

build:
	@go build -o ${BUILD_NAME} ./cmd/app

clean:
	@rm ${BUILD_NAME}

test:
	@go test ./internal/...
