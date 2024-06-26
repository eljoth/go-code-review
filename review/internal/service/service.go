package service

import (
	"coupon_service/internal/repository"
	"coupon_service/internal/service/entity"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) ApplyCoupon(basket entity.Basket, code string) (*entity.Basket, error) {
	if basket.Value < 0 {
		return nil, errors.New("tried to apply discount to negative value")
	}

	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, fmt.Errorf("error with coupon code %s: %w", code, err)
		}
		return nil, fmt.Errorf("failed to get coupon with code %s: %w", code, err)
	}

	if basket.Value > 0 {
		if basket.Value < coupon.MinBasketValue {
			basket.ApplicationSuccessful = false
			return &basket, nil
		}

		basket.AppliedDiscount = coupon.Discount
		basket.ApplicationSuccessful = true
	}

	return &basket, nil
}

func (s Service) CreateCoupon(discount int, code string, minBasketValue int) (*entity.Coupon, error) {
	coupon := entity.Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	if err := s.repo.Save(coupon); err != nil {
		return nil, fmt.Errorf("failed to save coupon: %w", err)
	}
	return &coupon, nil
}

func (s Service) GetCoupons(codes []string) ([]entity.Coupon, error) {
	coupons := make([]entity.Coupon, 0, len(codes))

	for _, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				return nil, fmt.Errorf("error with coupon code %s: %w", code, err)
			}
			return nil, fmt.Errorf("failed to get coupon with code %s: %w", code, err)
		}
		coupons = append(coupons, *coupon)
	}

	return coupons, nil
}
