FROM golang:1.16.0-alpine

WORKDIR /golang
ADD . /golang

ENV GO111MODULE=on

RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh \
    && go get github.com/gin-gonic/gin \
    && go get github.com/go-sql-driver/mysql \ 
    && go get github.com/jinzhu/gorm


EXPOSE 8080

CMD ["fresh"]
