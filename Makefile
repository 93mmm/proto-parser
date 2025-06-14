BUILD_NAME=protosym

build:
	@go build -o ${BUILD_NAME} ./cmd/app

clean:
	@rm ${BUILD_NAME}

test:
	@go test -count=1 ./internal/...
	@go test -count=1 ./tests/...

docker_build:
	@docker build . -t protosym -f internal/deploy/Dockerfile

# FIXME: путь к fila - не relative
docker_run: docker_build
	@docker run --rm -it protosym ./protosym /app/$(file)

docker_test: docker_build
	@docker run --rm -it protosym make test
