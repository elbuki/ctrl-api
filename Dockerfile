FROM golang:1.14.2

WORKDIR /go/src/app

COPY . .

RUN apk add --no-cache git musl-dev \
    && go get github.com/githubnemo/CompileDaemon \
    && apk del git

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon \
    --build="go build -o build/ ./src/..." \
    --command="./build/src"
