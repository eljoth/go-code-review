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

type Service interface {
	ApplyCoupon(*entity.Basket, string) error
	CreateCoupon(int, string, int) error
	GetCoupons([]string) ([]entity.Coupon, error)
}

type Config struct {
	Host string
	Port int
}

type API struct {
	srv *http.Server
	mux *gin.Engine
	svc Service
	cfg Config
}

func New(cfg Config, svc Service) API {
	r := gin.New()
	r.Use(gin.Recovery())

	return API{
		mux: r,
		cfg: cfg,
		svc: svc,
	}.withServer().withRoutes()
}

func (a API) withServer() API {
	a.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port),
		Handler: a.mux,
	}
	return a
}

func (a API) withRoutes() API {
	apiGroup := a.mux.Group("/api")
	apiGroup.POST("/apply", a.Apply)
	apiGroup.POST("/create", a.Create)
	apiGroup.GET("/coupons", a.Get)
	return a
}

func (a API) Start() {
	if err := a.srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (a API) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
