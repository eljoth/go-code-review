# 🚀 Coupon Service API

This is a simple coupon service API built with Go. It allows you to create, apply, and retrieve coupons.

## 🐳 Running with Docker

To run the service with Docker, follow these steps:

1. Build the Docker image:

```sh
docker build -t coupon-service .
```

2. Run the Docker container:

```sh
docker run -p 8080:8080 coupon-service
```

The service will be available at `http://localhost:8080`.


## Docker Hub
I have also uploaded the images to my [Docker Hub Profile](https://hub.docker.com/u/zvdy) Both for 32 and non 32 core CPU's.

## 📬 Making Requests

Here are some example `curl` commands to interact with the API:

- Create a coupon:

```sh
curl -X POST http://localhost:8080/api/create -d '{"discount": 10, "code": "Superdiscount", "minBasketValue": 50}' -H "Content-Type: application/json"
```

- Apply a coupon:

```sh
curl -X POST http://localhost:8080/api/apply -d '{"basket": {"value": 100}, "code": "Superdiscount"}' -H "Content-Type: application/json"
```

- Retrieve coupons: 

```sh
curl -X GET http://localhost:8080/api/coupons -d '{"codes": ["Superdiscount"]}' -H "Content-Type: application/json"
```

- Retrieve many coupons: 

```sh
curl -X GET http://localhost:8080/api/coupons -d '{"codes": ["Superdiscount1", "Superdiscount2", "Superdiscount3"]}' -H "Content-Type: application/json"
```

> You can use _[jq](https://jqlang.github.io/jq/_)_ in order to get formatted/prettier outputs just execute your curl command as usual, then add:  | jq and it will be formated 