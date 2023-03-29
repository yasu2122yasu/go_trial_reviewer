FROM golang:1.20-alpine3.16

COPY cmd /go/cmd/

RUN mkdir -p /usr/local/go/src/trial-4/cmd

WORKDIR /usr/local/go/src/trial-4/cmd

RUN apk update \
&& apk add --no-cache git \
&& go get github.com/gin-gonic/gin \
&& go get github.com/jinzhu/gorm \
&& go get github.com/lib/pq \
# ここ不安あるので後で確認する。
&& go mod init trial4 \
&& apk add git