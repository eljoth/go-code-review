# pin version
# strip debug symbols
# static linked binary should work

FROM golang:alpine AS builder
COPY . /go/src/coupon-service
WORKDIR /go/src/coupon-service/cmd/coupon_service
RUN go build -o coupon-service .

FROM alpine:latest
COPY --from=builder /go/src/coupon-service/cmd/coupon_service/coupon-service /coupon-service
ENV GIN_MODE=release \
    API_HOST=0.0.0.0
ENTRYPOINT /coupon-service
EXPOSE 8080
