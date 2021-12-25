# Earthfile
VERSION 0.6
FROM golang:1.16-alpine3.13
RUN apk --update --no-cache add git curl
WORKDIR /app

all:
  BUILD +lint
  BUILD +docker

build:
  COPY . .
  RUN CGO_ENABLED=0 go build -o build/ep main.go
  SAVE ARTIFACT build/ep AS LOCAL build/ep

lint:
  RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
  COPY . .
  RUN CGO_ENABLED=0 golangci-lint run --timeout 5m

docker:
  COPY +build/ep .
  RUN mv /app/ep /usr/local/bin/ep && rm -rf /app
  SAVE IMAGE ep:latest
