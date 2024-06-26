package api

import (
	"context"
	"coupon_service/internal/service/entity"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Service interface {
	ApplyCoupon(entity.Basket, string) (*entity.Basket, error)
	CreateCoupon(int, string, int) (*entity.Coupon, error)
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

func New[T Service](cfg Config, svc T) *API {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	api := &API{
		mux: r,
		cfg: cfg,
		svc: svc,
	}

	return api.withRoutes().withServer()
}

func (a *API) withServer() *API {
	a.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port),
		Handler: a.mux,
	}

	return a
}

func (a *API) withRoutes() *API {
	apiGroup := a.mux.Group("/api")
	{
		apiGroup.POST("/apply", a.Apply)
		apiGroup.POST("/create", a.Create)
		apiGroup.GET("/coupons", a.Get)
	}

	return a
}

func (a *API) Start() error {
	if err := a.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server terminated unexpectedly: %w", err)
	}
	return nil
}

func (a *API) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}
	return nil
}
