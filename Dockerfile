# build stage
FROM golang:1.22-alpine3.18 AS builder

WORKDIR /go/src/go-code-review

COPY . .

RUN apk add --no-cache --virtual git

RUN go build -o ./cmd/server/main.go

ENTRYPOINT ./main
EXPOSE 8080
