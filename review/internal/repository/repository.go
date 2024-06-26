package repository

import (
	"coupon_service/internal/service/entity"

	"github.com/pkg/errors"
)

type Repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(entity.Coupon) error
}

var ErrNotFound = errors.New("coupon not found")
