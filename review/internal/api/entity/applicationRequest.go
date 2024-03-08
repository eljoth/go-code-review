package entity

import svcentity "coupon_service/internal/service/entity"

type ApplicationRequest struct {
	Code   string           `binding:"required"`
	Basket svcentity.Basket `binding:"required"`
}
