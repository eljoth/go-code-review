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

func (s Service) ApplyCoupon(b *Basket, code string) error {
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return err
	}

	if b.Value < coupon.MinBasketValue {
		return fmt.Errorf("basket value too low for this code")
	}

	b.Value -= coupon.Discount
	b.ApplicationSuccessful = true

	if b.Value < 0 {
		b.AppliedDiscount = coupon.Discount + b.Value
		b.Value = 0
	} else {
		b.AppliedDiscount = coupon.Discount
	}
	return nil
}

func (s Service) CreateCoupon(discount int, code string, minBasketValue int) error {
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
		} else {
			coupons = append(coupons, *coupon)
		}
	}

	return coupons, e
}
