package memdb

import (
	"coupon_service/internal/service/entity"
	"fmt"
)

// Config is a struct that can be used to configure the repository.
type Config struct{}

// repository is an interface that defines the methods that a coupon repository should have.
type repository interface {
	// FindByCode finds a coupon by its code.
	FindByCode(string) (*entity.Coupon, error)
	// Save saves a coupon to the repository.
	Save(entity.Coupon) error
}

// Repository is a struct that holds the in-memory representation of the coupon repository.
type Repository struct {
	// entries is a map that stores the coupons, using their code as the key.
	entries map[string]entity.Coupon
}

// New creates a new repository and returns a pointer to it.
func New() *Repository {
	return &Repository{
		// Initialize the entries map.
		entries: make(map[string]entity.Coupon),
	}
}

// FindByCode is a method that finds a coupon by its code in the repository.
func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
	// Try to find the coupon in the entries map.
	coupon, ok := r.entries[code]
	// If the coupon is not found, return an error.
	if !ok {
		return nil, fmt.Errorf("coupon not found")
	}
	// If the coupon is found, return a pointer to it.
	return &coupon, nil
}

// Save is a method that saves a coupon to the repository.
func (r *Repository) Save(coupon entity.Coupon) error {
	// Save the coupon to the entries map, using its code as the key.
	r.entries[coupon.Code] = coupon
	// Return nil as there is no error.
	return nil
}
