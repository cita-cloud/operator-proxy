FROM golang:1.17 as builder

WORKDIR /workspace
# ENV GOPROXY https://goproxy.cn
# Copy the Go Modules manifests
# # cache deps before building and copying source so that we don't need to re-download as much
# # and so that source changes don't invalidate our downloaded layer
#
# # Copy the go source
ENV GOPROXY https://goproxy.cn,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
#
# # Build
COPY api/ api/
COPY server/ server/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o operator-proxy server/main.go
#
# # Use distroless as minimal base image to package the manager binary
# # Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

ARG version
ENV GIT_COMMIT=$version
LABEL maintainers="rivtower.com"
LABEL description="operator-proxy"
MAINTAINER https://github.com/acechef

WORKDIR /
COPY --from=builder /workspace/operator-proxy .
USER 65532:65532
#
ENTRYPOINT ["/operator-proxy"]
