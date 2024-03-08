package api

import (
	. "coupon_service/internal/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a API) Apply(c *gin.Context) {
	apiReq := ApplicationRequest{}
	if err := c.BindJSON(&apiReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := a.svc.ApplyCoupon(&apiReq.Basket, apiReq.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiReq.Basket)
		return
	}
	c.JSON(http.StatusOK, apiReq.Basket)
}

func (a API) Create(c *gin.Context) {
	apiReq := Coupon{}
	if err := c.BindJSON(&apiReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		_ = c.AbortWithError(http.StatusConflict, err)
		return
	}
}

func (a API) Get(c *gin.Context) {
	apiReq := CouponRequest{}
	if err := c.BindJSON(&apiReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	coupons, _ := a.svc.GetCoupons(apiReq.Codes)
	c.JSON(http.StatusOK, coupons)
}
