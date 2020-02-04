FROM golang:1.13.7-alpine3.11

ENV GOBIN /go/bin
RUN apk add --no-cache git

RUN go get -u github.com/golang/dep/cmd/dep && go get -u github.com/githubnemo/CompileDaemon
VOLUME ["/go/src/github.com/karidoe/app"]