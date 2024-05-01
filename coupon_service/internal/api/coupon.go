package api

import (
	//changed dot import https://go.dev/wiki/CodeReviewComments#import-dot and added entity. prefix
	"coupon_service/internal/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Apply is a handler function that applies a coupon to a basket
func (a *API) Apply(c *gin.Context) {
	apiReq := entity.ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, basket)
}

// Create is a handler function that creates a new coupon
func (a *API) Create(c *gin.Context) {
	apiReq := entity.Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if a coupon with the same code already exists
	_, err := a.svc.FindByCode(apiReq.Code)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "A coupon with this code already exists"})
		return
	}

	coupon, err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, coupon) // Return Created 201 with the coupon details
}

// Get is a handler function that retrieves a list of coupons based on the provided codes
func (a *API) Get(c *gin.Context) {
	apiReq := entity.CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coupons)
}
