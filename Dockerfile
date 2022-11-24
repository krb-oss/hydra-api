FROM golang:1.19.2-bullseye AS builder
WORKDIR /build

ENV DEBIAN_FRONTEND noninteractive
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

COPY . .
RUN go build -ldflags="-w -s" -o /main ./cmd/worker

RUN ap-get update && apt-get install upx --yes
RUN upx /main

FROM scratch
COPY --from=builder /main /main
ENTRYPOINT ["/main"]
