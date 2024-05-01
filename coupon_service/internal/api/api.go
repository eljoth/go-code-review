package api

import (
	"context"
	"coupon_service/internal/service/entity"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Service interface defines the methods that our service should implement.
type Service interface {
	ApplyCoupon(entity.Basket, string) (*entity.Basket, error)
	CreateCoupon(int, string, int) (*entity.Coupon, error)
	GetCoupons([]string) ([]entity.Coupon, error)
	FindByCode(string) (*entity.Coupon, error)
}

// Config struct holds the configuration for our API.
type Config struct {
	Host string
	Port int
}

// API struct holds the server, router, service and config for our API.
type API struct {
	srv *http.Server
	MUX *gin.Engine
	svc Service
	CFG Config
}

// New function creates a new API with the given config and service.
func New[T Service](cfg Config, svc T) API {
	gin.SetMode(gin.ReleaseMode)
	r := new(gin.Engine)
	r = gin.New()
	r.Use(gin.Recovery())

	return API{
		MUX: r,
		CFG: cfg,
		svc: svc,
	}.withServer()
}

// withServer function creates a new server with the given address and handler.
func (a API) withServer() API {

	ch := make(chan API)
	go func() {
		a.srv = &http.Server{
			Addr:    fmt.Sprintf(":%d", a.CFG.Port),
			Handler: a.MUX,
		}
		ch <- a
	}()

	return <-ch
}

// withRoutes function adds routes to the API.
func (a API) withRoutes() API {
	apiGroup := a.MUX.Group("/api")
	apiGroup.Use(gin.Logger()) // Log all requests

	apiGroup.POST("/apply", a.Apply)
	apiGroup.POST("/create", a.Create)
	apiGroup.GET("/coupons", a.Get)
	return a
}

/* Curl commands to test the API

jq is helpful for pretty printing the JSON responses, not needed

curl -X POST http://localhost:8080/api/create -d '{"discount": 10, "code": "Superdiscount", "minBasketValue": 50}' -H "Content-Type: application/json" | jq
curl -X POST http://localhost:8080/api/apply -d '{"basket": {"value": 100}, "code": "Superdiscount"}' -H "Content-Type: application/json" | jq
curl -X GET http://localhost:8080/api/coupons -d '{"codes": ["Superdiscount"]}' -H "Content-Type: application/json" | jq

*/

// Start function starts the server and serves the API routes.
func (a API) Start() {
	go func() {
		fmt.Println("Server starting...")
		if err := a.srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Serve the API routes
	a.withRoutes()
}

// Close function shuts down the server after a delay.
func (a API) Close() {
	<-time.After(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
