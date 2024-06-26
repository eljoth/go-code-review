# build stage
FROM golang:1.22-alpine3.20 AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /go/src/coupon-service
COPY . .
RUN go build -o ./build/coupon_service ./cmd/coupon_service/main.go

# final stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /go/src/coupon-service/build/coupon_service /usr/bin/coupon_service

ENTRYPOINT ["/usr/bin/coupon_service"]
