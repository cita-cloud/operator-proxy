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
	docker build --platform linux/arm64 --network=host -t $(PROTOC_IMAGE_NAME):$(PROTOC_IMAGE_VERSION) -f ./Dockerfile-protoc-3.19.1 .

protoc-image-push:
	docker push $(PROTOC_IMAGE_NAME):$(PROTOC_IMAGE_VERSION)

grpc-code-generate:
	docker run -v $(PWD):/src -e GO111MODULE=on $(PROTOC_IMAGE_NAME):$(PROTOC_IMAGE_VERSION) /bin/bash ./grpc-code-generate.sh

LD_FLAGS=-ldflags " \
    -X $(shell go list -m)/cli/command.ClientVersion=0.0.1 \
    -X $(shell go list -m)/cli/command.Goos=$(shell go env GOOS) \
    -X $(shell go list -m)/cli/command.Goarch=$(shell go env GOARCH) \
    -X $(shell go list -m)/cli/command.GitCommit=$(shell git rev-parse HEAD) \
    -X $(shell go list -m)/cli/command.BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ') \
    "

build: fmt linux-amd-cli linux-arm-cli mac-amd-cli mac-arm-cli win-cli

linux-amd-cli: GO_ENV += GOOS=linux GOARCH=amd64
linux-amd-cli:
	$(GO_BUILD) $(LD_FLAGS) -o bin/cco-cli ./cli

linux-arm-cli: GO_ENV += GOOS=linux GOARCH=arm64
linux-arm-cli:
	$(GO_BUILD) $(LD_FLAGS) -o bin/cco-cli ./cli

mac-amd-cli: GO_ENV += GOOS=darwin GOARCH=amd64
mac-amd-cli:
	$(GO_BUILD) $(LD_FLAGS) -o bin/cco-cli ./cli

mac-arm-cli: GO_ENV += GOOS=darwin GOARCH=arm64
mac-arm-cli:
	$(GO_BUILD) $(LD_FLAGS) -o bin/cco-cli ./cli

win-cli: GO_ENV += GOOS=windows GOARCH=386
win-cli:
	$(GO_BUILD) $(LD_FLAGS) -o bin/cco-cli.exe ./cli

docker-build: ## Build docker image with the manager.
	docker build --platform linux/amd64 -t ${IMG} .

docker-push: ## Push docker image with the manager.
	docker push ${IMG}