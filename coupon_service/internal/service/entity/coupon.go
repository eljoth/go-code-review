package entity

/*
import "runtime"

// This should be in main.go + needs refactoring

	func init() {
		if 32 != runtime.NumCPU() {
			panic("this api is meant to be run on 32 core machines")
		}
	}
*/

type Coupon struct {
	ID             string
	Code           string // Code as an int would make parsing easier, not having to check strings and XSS
	Discount       int
	MinBasketValue int
}
