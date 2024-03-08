package entity

type Basket struct {
	Value                 int  `binding:"required"`
	AppliedDiscount       int  `binding:"-"`
	ApplicationSuccessful bool `binding:"-"`
}
