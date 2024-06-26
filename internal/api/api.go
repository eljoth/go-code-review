package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"coupon_service/internal/service/entity"

	"github.com/gin-gonic/gin"
)

type Service interface {
	ApplyCoupon(basketValue int, code string) (*entity.Basket, error)
	CreateCoupon(code string, discount, minBasketValue int)
	GetCoupons(codes []string) ([]*entity.Coupon, error)
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

func New[T Service](cfg Config, svc T) API {
	gin.SetMode(gin.ReleaseMode)
	r := new(gin.Engine)
	r = gin.New()
	r.Use(gin.Recovery())

	return API{
		mux: r,
		cfg: cfg,
		svc: svc,
	}.withServer()
}

func (a API) withServer() API {

	ch := make(chan API)
	go func() {
		a.srv = &http.Server{
			Addr:    fmt.Sprintf(":%d", a.cfg.Port),
			Handler: a.mux,
		}
		ch <- a
	}()

	return <-ch
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
		panic(err)
	}
}

func (a API) Close() {
	<-time.After(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
