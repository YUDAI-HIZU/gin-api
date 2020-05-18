FROM golang:1.12.0-alpine3.9

WORKDIR /go/src/github.com/YUDAI-HIZU/gin-api/server

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]