package api

import (
	"coupon_service/internal/repository"
	"coupon_service/internal/service/entity"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApplicationRequest struct {
	Code   string `json:"code"`
	Basket Basket `json:"basket"`
}

func (a ApplicationRequest) validate() error {
	if a.Code == "" {
		return errors.New("code must not be empty")
	}
	return a.Basket.validate()
}

type Basket struct {
	Value                 int  `json:"value"`
	AppliedDiscount       int  `json:"applied_discount"`
	ApplicationSuccessful bool `json:"application_successful"`
}

func (b Basket) validate() error {
	if b.Value < 0 {
		return errors.New("value must be greater than or equal to 0")
	}
	return nil
}

func (b Basket) toEntity() entity.Basket {
	return entity.Basket{
		Value:                 b.Value,
		AppliedDiscount:       b.AppliedDiscount,
		ApplicationSuccessful: b.ApplicationSuccessful,
	}
}

func basketFromEntity(e entity.Basket) Basket {
	return Basket{
		Value:                 e.Value,
		AppliedDiscount:       e.AppliedDiscount,
		ApplicationSuccessful: e.ApplicationSuccessful,
	}
}

func (a *API) Apply(c *gin.Context) {
	var apiReq ApplicationRequest
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	if err := apiReq.validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	basket, err := a.svc.ApplyCoupon(apiReq.Basket.toEntity(), apiReq.Code)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, basketFromEntity(*basket))
}

type Coupon struct {
	Discount       int    `json:"discount"`
	Code           string `json:"code"`
	MinBasketValue int    `json:"min_basket_value"`
}

func (c Coupon) validate() error {
	if c.Discount < 1 {
		return errors.New("discount must be greater than 0")
	}
	if c.MinBasketValue < 0 {
		return errors.New("min_basket_value must be greater than or equal to 0")
	}
	if c.Code == "" {
		return errors.New("code must not be empty")
	}
	return nil
}

func couponFromEntity(e entity.Coupon) Coupon {
	return Coupon{
		Discount:       e.Discount,
		Code:           e.Code,
		MinBasketValue: e.MinBasketValue,
	}
}

func (a *API) Create(c *gin.Context) {
	var apiReq Coupon
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request"})
		return
	}
	if err := apiReq.validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coupon, err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		return
	}
	c.JSON(http.StatusCreated, couponFromEntity(*coupon))
}

func (a *API) Get(c *gin.Context) {
	codes, ok := c.GetQueryArray("codes")
	if !ok || len(codes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "codes must not be empty"})
		return
	}

	coupons, err := a.svc.GetCoupons(codes)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coupons)
}
