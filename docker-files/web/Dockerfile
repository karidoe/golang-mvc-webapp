FROM golang:1.13.7-alpine3.11

RUN apk add --update --no-cache git \
    && apk add --no-cache gcc libc-dev librdkafka-dev pkgconf

RUN go get -u github.com/golang/dep/cmd/dep \
    && go get -u github.com/githubnemo/CompileDaemon