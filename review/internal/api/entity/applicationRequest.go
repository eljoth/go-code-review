package entity

import "coupon_service/review/internal/service/entity"

type ApplicationRequest struct {
	Code   string
	Basket entity.Basket
}
