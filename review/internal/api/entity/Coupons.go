package entity

type Coupon struct {
	Discount       int    `binding:"required"`
	Code           string `binding:"required"`
	MinBasketValue int    `binding:"required"`
}
