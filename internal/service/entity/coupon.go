package entity

import (
	"github.com/google/uuid"
)

type Coupon struct {
	ID             uuid.UUID `json:"id"`
	Code           string    `json:"code"`
	Discount       int       `json:"discount"`
	MinBasketValue int       `json:"min_basket_value"`
}
