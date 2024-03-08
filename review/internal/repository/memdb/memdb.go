package memdb

import (
	"coupon_service/internal/service/entity"
	"fmt"
)

type Repository struct {
	entries map[string]entity.Coupon
}

func New() *Repository {
	return &Repository{
		entries: make(map[string]entity.Coupon),
	}
}

func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
	coupon, ok := r.entries[code]
	if !ok {
		return nil, fmt.Errorf("Coupon for code %s not found", code)
	}
	return &coupon, nil
}

func (r *Repository) Save(coupon entity.Coupon) error {
	if _, dup := r.entries[coupon.Code]; dup {
		return fmt.Errorf("coupon for code %s already exists, wont overwrite", coupon.Code)
	}
	r.entries[coupon.Code] = coupon
	return nil
}
