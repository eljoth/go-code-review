package service

import (
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service/entity"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		repo Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"initialize service", args{repo: nil}, Service{repo: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ApplyCoupon(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		basket entity.Basket
		code   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantB   *entity.Basket
		wantErr bool
	}{
		{"Apply 10%", fields{memdb.New()}, args{entity.Basket{Value: 16}, "DiscountCode"}, &entity.Basket{Value: 14.4, AppliedDiscount: 10, ApplicationSuccessful: true}, false},
		{"Apply 10%", fields{memdb.New()}, args{entity.Basket{Value: 14}, "DiscountCode"}, nil, true},
		{"Apply 10%", fields{memdb.New()}, args{entity.Basket{Value: 0}, "DiscountCode"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}
			s.CreateCoupon(10, tt.args.code, 15)
			gotB, err := s.ApplyCoupon(tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ApplyCoupon() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestService_CreateCoupon(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{"Apply 10%", fields{memdb.New()}, args{10, "Superdiscount", 55}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}

			s.CreateCoupon(tt.args.discount, tt.args.code, tt.args.minBasketValue)
		})
	}
}

func TestService_GetCoupons(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Apply 10%", fields{memdb.New()}, args{10, "Superdiscount", 55}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}

			s.CreateCoupon(tt.args.discount, tt.args.code, tt.args.minBasketValue)
			coupons, err := s.GetCoupons([]string{tt.args.code})
			if len(coupons) > 1 {
				t.Errorf("GetCoupons() should return one coupon")
			}
			if err != nil {
				t.Errorf("ApplyCoupon() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
