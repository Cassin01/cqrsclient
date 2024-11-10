FROM golang:1.23.3-alpine3.20
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/client
WORKDIR /go/src/client
ADD . /go/src/client
