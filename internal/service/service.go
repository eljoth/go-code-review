package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"coupon_service/internal/repository/memdb"
	. "coupon_service/internal/service/entity"
)

type Service struct {
	memDBRepo memdb.MemDB
}

func New(memDBRepo memdb.MemDB) (*Service, error) {
	if memDBRepo == nil {
		return nil, errors.New("repo cannot be nil")
	}
	return &Service{memDBRepo: memDBRepo}, nil
}

func (s *Service) ApplyCoupon(basketValue int, code string) (*Basket, error) {
	if basketValue <= 0 {
		return nil, errors.New("basketValue value cannot be less or equal to zero")
	}

	coupon, err := s.memDBRepo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	return &Basket{
		Value:                 basketValue,
		AppliedDiscount:       coupon.Discount,
		ApplicationSuccessful: true,
	}, nil
}

func (s *Service) CreateCoupon(code string, discount, minBasketValue int) {
	s.memDBRepo.Save(&Coupon{
		ID:             uuid.New(),
		Code:           code,
		Discount:       discount,
		MinBasketValue: minBasketValue,
	})
}

func (s *Service) GetCoupons(codes []string) ([]*Coupon, error) {
	coupons := make([]*Coupon, 0, len(codes))
	var e error

	for _, code := range codes {
		coupon, err := s.memDBRepo.FindByCode(code)
		if err != nil {
			e = fmt.Errorf("%w; code: %s", e, code)
		}
		coupons = append(coupons, coupon)
	}

	return coupons, e
}
