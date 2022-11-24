VERSION 0.6

FROM golang:1.19.3-alpine

RUN apk add --update --no-cache \
    coreutils \
    g++ \
    git \
    make

WORKDIR /app

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

sources:
    FROM +deps
    COPY --dir ./*.go cmd postgres ./

build:
    FROM +sources
    RUN go build -ldflags="-w -s" -o tmp/main ./cmd
    SAVE ARTIFACT tmp/main AS LOCAL tmp/main

lint:
    FROM +sources
    RUN go install github.com/mgechev/revive@latest
    COPY revive.toml ./
    RUN revive -config revive.toml ./...

unit-test:
    FROM +sources
    RUN go test ./...

docker:
    COPY +build/main .
    ENTRYPOINT ["/app/main"]
    SAVE IMAGE karlbateman/hydro-api:latest

all:
    BUILD +lint
    BUILD +build
    BUILD +unit-test
    BUILD +docker
