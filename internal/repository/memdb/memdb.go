package memdb

import (
	"errors"

	"github.com/google/uuid"

	"coupon_service/internal/service/entity"
)

type MemDB interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(*entity.Coupon)
}

var _ MemDB = (*Repository)(nil) // we ensure that we comply with the interface

type Repository struct {
	couponEntries map[string]*entity.Coupon
}

func New() *Repository {
	return &Repository{
		couponEntries: map[string]*entity.Coupon{
			"ES0001": {
				ID:             uuid.New(),
				Code:           "ES0001",
				Discount:       10,
				MinBasketValue: 1,
			},
			"ES0002": {
				ID:             uuid.New(),
				Code:           "ES0002",
				Discount:       10,
				MinBasketValue: 1,
			},
		},
	}
}

func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
	coupon, ok := r.couponEntries[code]
	if !ok {
		return nil, errors.New("coupon not found")
	}
	return coupon, nil
}

func (r *Repository) Save(coupon *entity.Coupon) {
	r.couponEntries[coupon.Code] = coupon
}
