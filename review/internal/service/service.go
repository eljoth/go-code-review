package service

import (
	. "coupon_service/internal/service/entity"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	FindByCode(string) (*Coupon, error)
	Save(Coupon) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) ApplyCoupon(basket Basket, code string) (b *Basket, e error) {
	b = &basket
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if coupon.MinBasketValue > int(b.Value) {
		return nil, fmt.Errorf("basket value does not reach minimum value")
	}

	if b.Value > 0 {
		b.Value = b.Value * (1 - float64(coupon.Discount)/100)
		b.AppliedDiscount = coupon.Discount
		b.ApplicationSuccessful = true
		return b, nil
	}

	return nil, fmt.Errorf("discount value cannot be zero")
}

func (s Service) CreateCoupon(discount int, code string, minBasketValue int) any {
	coupon := Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	if err := s.repo.Save(coupon); err != nil {
		return err
	}
	return nil
}

func (s Service) GetCoupons(codes []string) ([]Coupon, error) {
	coupons := make([]Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			if e == nil {
				e = fmt.Errorf("code: %s, index: %d", code, idx)
			} else {
				e = fmt.Errorf("%w; code: %s, index: %d", e, code, idx)
			}
			continue
		}
		coupons = append(coupons, *coupon)
	}

	return coupons, e
}
