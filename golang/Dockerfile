FROM golang:1.16.0-alpine

WORKDIR /go/src/app
ADD . /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh \
    && go get github.com/gin-gonic/gin


EXPOSE 8080

CMD ["fresh"]
