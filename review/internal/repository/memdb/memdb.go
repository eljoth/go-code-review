package memdb

import (
	"coupon_service/internal/repository"
	"coupon_service/internal/service/entity"
	"sync"
)

var ()

type Repository struct {
	entries map[string]entity.Coupon
	lock    sync.RWMutex
}

func New() *Repository {
	return &Repository{
		entries: make(map[string]entity.Coupon),
	}
}

func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	coupon, ok := r.entries[code]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &coupon, nil
}

func (r *Repository) Save(coupon entity.Coupon) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.entries[coupon.Code] = coupon
	return nil
}
