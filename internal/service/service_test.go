package service_test

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"

	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"coupon_service/internal/service/entity"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		repo    memdb.MemDB
		wantErr error
	}{
		{
			name:    "failed to init service",
			repo:    nil,
			wantErr: errors.New("repo cannot be nil"),
		},
		{
			name:    "ok",
			repo:    memdb.New(),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.New(tt.repo)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestService_ApplyCoupon(t *testing.T) {
	repo := memdb.New()
	s, err := service.New(repo)
	if err != nil {
		panic(err)
	}

	type args struct {
		basketValue int
		code        string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Basket
		wantErr error
	}{
		{
			name: "ok",
			args: args{
				basketValue: 100,
				code:        "ES0001",
			},
			want: &entity.Basket{
				Value:                 100,
				AppliedDiscount:       10,
				ApplicationSuccessful: true,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := s.ApplyCoupon(tt.args.basketValue, tt.args.code)
			assert.Equal(t, tt.want, res)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestService_CreateCoupon(t *testing.T) {
	repo := memdb.New()
	s, err := service.New(repo)
	if err != nil {
		panic(err)
	}

	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name   string
		args   args
		assert func()
	}{
		{
			name: "ok",
			args: args{10, "ES0001", 55},
			assert: func() {
				_, err = repo.FindByCode("ES0001")
				assert.Equal(t, nil, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.CreateCoupon(tt.args.code, tt.args.discount, tt.args.minBasketValue)
			if tt.assert != nil {
				tt.assert()
			}
		})
	}
}

func TestService_GetCoupons(t *testing.T) {
	repo := memdb.New()
	s, err := service.New(repo)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		codes   []string
		want    []*entity.Coupon
		wantErr error
	}{
		{
			name:  "ok",
			codes: []string{"ES0001", "error-code"},
			want: []*entity.Coupon{
				{
					Code:           "ES0001",
					Discount:       10,
					MinBasketValue: 1,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := s.GetCoupons(tt.codes)
			assert.Equal(t, tt.want, res)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
