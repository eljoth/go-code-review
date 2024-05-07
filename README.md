# Schwarz IT Code Review Repository

API SERVICE for managing discount coupons. 

# Starting the project 

- Without docker: 

Just go in review/cmd/coupon_service and run `go run main.go`

- Using docker:

1. Build the docker image:

```sh
docker build -t coupon_api_service .
```

2. Run it 

```sh
docker run -p 8080:8080 coupon-service
```

Voila! You can see the server running at `http://localhost:8080`

## API

Below is a list of API endpoints with their respective input and output.

### Creating coupons

#### Endpoint

```
POST
/api/coupon/create
```

#### Input

```json
{
    "Discount": <int>,
    "Code": <string>,
    "MinBasketValue": <int>
}
```

### Listing coupons

#### Endpoint

```
GET
/api/coupon/list
```

#### Input

```json
{
  "Codes": [<string>, <string>, <string>]
}
```

### Apply coupon value

#### Endpoint

```
POST
/api/coupon/apply
```

#### Input

```json
{
  "Code": <string>,
  "Basket": {
    "Value"                 <int>
	"AppliedDiscount"       <int>
	"ApplicationSuccessful" <bool>
  }
}
```