APP?=dashboardserver
COMMIT_SHA=$(shell git rev-parse --short HEAD)
DOCKER_REGISTRY?=docker.io


.PHONY: build
## build: the application
build: clean
	@echo "Building..."
	@go build -o ${APP} cmd/api/main.go

.PHONY: run
## run: runs go run main.go
run:
	go run -race cmd/api/main.go

.PHONY: clean
## clean: cleans the binary
	@echo "Cleaning"
	@go clean

.PHONY: test
## test: runs go tests with default values
test:
	go test -v -cover -count=1 -race ./...

.PHONY: docker-login
## docker-login: log in to a Docker registry
docker-login:
	docker login -u ${DOCKER_USER} -p ${DOCKER_PASSWORD}

.PHONY: podman-login
## podman-login: log in to a Docker registry
podman-login:
	podman login -u ${DOCKER_USER} -p ${DOCKER_PASSWORD} ${DOCKER_REGISTRY}

.PHONY: docker-build
## docker-build: builds the dashboardserver docker image to registry
docker-build: build
	docker build -t ${APP}:${COMMIT_SHA} .
	docker tag ${APP}:${COMMIT_SHA} ${DOCKER_USER}/${APP}:${COMMIT_SHA}

.PHONY: podman-build
## podman-build: builds the dashboardserver podman image
podman-build: build
	podman build -t ${APP}:${COMMIT_SHA} .
	podman tag ${APP}:${COMMIT_SHA} ${DOCKER_USER}/${APP}:${COMMIT_SHA}

.PHONY: docker-push
## docker-push: pushes the dashboardserver docker image to registry
docker-push: docker-login docker-build
	docker push ${DOCKER_USER}/${APP}:${COMMIT_SHA}

.PHONY: podman-push
## podman-push: pushes the dashboardserver image to registry
podman-push: podman-login podman-build
	podman push ${DOCKER_USER}/${APP}:${COMMIT_SHA}

.PHONY: up
## up: builds and starts containers for a service
up:
	docker-compose up --build --detach

.PHONY: down
## down: stops containers and remove containers, networks, volumes and images created by up
down:
	docker-compose down

.PHONY: logs
## logs: Displays log output from services
logs:
	docker-compose logs --follow

.PHONY: podman-up
## podman-up: builds and starts containers for a service
podman-up:
	podman-compose up --build --detach

.PHONY: podman-down
## podman-down: stops containers and remove containers, networks, volumes and images created by up
podman-down:
	podman-compose down

.PHONY: check_install
## check_install: check if swagger is installed
check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: swag
## swagger: generate or update the swagger config file
swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
