# build stage
FROM golang:latest AS builder

RUN apt-get update && \
    apt-get -y install gcc libc-dev 


WORKDIR /go/src

COPY . .

WORKDIR /go/src
RUN go build -o main ./cmd/coupon_service/main.go

ENTRYPOINT ./main
EXPOSE 8080
