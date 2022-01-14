GO_BUILD = $(GO_CMD) build
GO_CMD = $(GO_ENV) go

#IMG ?= citacloud/operator-proxy:v0.0.1
IMG ?= registry.devops.rivtower.com/cita-cloud/operator/operator-proxy:v0.0.1

PROTOC_IMAGE_NAME=registry.devops.rivtower.com/cita-cloud/operator/protoc
PROTOC_IMAGE_VERSION=3.19.1

# Run go fmt against code
fmt:
	go fmt ./...

protoc-image-build:
	docker build --network=host -t $(PROTOC_IMAGE_NAME):$(PROTOC_IMAGE_VERSION) -f ./Dockerfile-protoc-3.19.1 .
	#docker build -t $(PROTOC_IMAGE_NAME):$(PROTOC_IMAGE_VERSION) -f ./Dockerfile-protoc-3.19.1 .

protoc-image-push:
	docker push $(PROTOC_IMAGE_NAME):$(PROTOC_IMAGE_VERSION)

grpc-code-generate:
	docker run -v $(PWD):/src -e GO111MODULE=on $(PROTOC_IMAGE_NAME):$(PROTOC_IMAGE_VERSION) /bin/bash ./grpc-code-generate.sh


build: fmt mac-cli

mac-cli: GO_ENV += GOOS=darwin GOARCH=amd64
mac-cli:
	$(GO_BUILD) -o bin/cco-cli-mac ./cli

docker-build: ## Build docker image with the manager.
	docker build -t ${IMG} .

docker-push: ## Push docker image with the manager.
	docker push ${IMG}