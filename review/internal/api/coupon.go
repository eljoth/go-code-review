package api

import (
	"coupon_service/internal/service/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApplicationRequest struct {
	Code   string
	Basket entity.Basket
}

func (a *API) Apply(c *gin.Context) {
	apiReq := ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, basket)
}

type Coupon struct {
	Discount       int
	Code           string
	MinBasketValue int
}

func (a *API) Create(c *gin.Context) {
	apiReq := Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		return
	}
	c.Status(http.StatusOK)
}

type CouponRequest struct {
	Codes []string
}

func (a *API) Get(c *gin.Context) {
	apiReq := CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, coupons)
}
