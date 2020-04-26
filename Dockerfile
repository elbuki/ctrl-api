FROM golang:1.14.2-alpine

WORKDIR /go/src/app

COPY . .

RUN apk add --no-cache git \
    && go get github.com/githubnemo/CompileDaemon \
    && apk del git

ENTRYPOINT CompileDaemon \
    --build="go build -o build/ ./src/..." \
    --command="./build/src"
