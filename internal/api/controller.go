package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"coupon_service/internal/service/entity"
)

func (a API) Apply(c *gin.Context) {
	type ApplicationRequest struct {
		Code        string `json:"code"`
		BasketValue int    `json:"basket_value"`
	}

	req := ApplicationRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	basket, err := a.svc.ApplyCoupon(req.BasketValue, req.Code)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, basket)
}

func (a API) Create(c *gin.Context) {
	req := entity.Coupon{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	a.svc.CreateCoupon(req.Code, req.Discount, req.MinBasketValue)

	c.Status(http.StatusNoContent)
}

func (a API) Get(c *gin.Context) {
	type CodesRequest struct {
		Codes []string `json:"codes"`
	}
	type CouponsResponse struct {
		Coupons []*entity.Coupon `json:"coupons"`
		Errors  error
	}

	req := CodesRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	coupons, err := a.svc.GetCoupons(req.Codes)

	c.JSON(http.StatusOK, CouponsResponse{
		Coupons: coupons,
		Errors:  err,
	})
}
