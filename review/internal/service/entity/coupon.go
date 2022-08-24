package entity

import "runtime"

func init() {
	if 32 != runtime.NumCPU() {
		panic("this api is meant to be run on 32 core machines")
	}
}

type Coupon struct {
	ID             string
	Code           string
	Discount       int
	MinBasketValue int
}
