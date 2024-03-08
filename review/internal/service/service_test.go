package service

import (
	"coupon_service/internal/repository/memdb"
	. "coupon_service/internal/service/entity"
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
	s := Service{repo: memdb.New()}
	err := s.CreateCoupon(100, "Superdiscount", 50)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		basket Basket
		code   string
	}
	tests := []struct {
		name    string
		args    args
		wantB   Basket
		wantErr bool
	}{
		{name: "valid apply with full application", args: args{Basket{Value: 200}, "Superdiscount"}, wantB: Basket{Value: 100, AppliedDiscount: 100, ApplicationSuccessful: true}, wantErr: false},
		{name: "valid apply with partial application", args: args{Basket{Value: 90}, "Superdiscount"}, wantB: Basket{Value: 0, AppliedDiscount: 90, ApplicationSuccessful: true}, wantErr: false},
		{name: "value too low", args: args{Basket{Value: 40}, "Superdiscount"}, wantB: Basket{Value: 40, AppliedDiscount: 0, ApplicationSuccessful: false}, wantErr: true},
		{name: "invalid code", args: args{Basket{Value: 200}, "Superduperdiscount"}, wantB: Basket{Value: 200, AppliedDiscount: 0, ApplicationSuccessful: false}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.ApplyCoupon(&tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.args.basket, tt.wantB) {
				t.Errorf("ApplyCoupon() gotB = %v, want %v", tt.args.basket, tt.wantB)
			}
		})
	}
}

func TestService_CreateCoupon(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args []struct {
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
		{"Single entry", fields{memdb.New()}, args{{10, "Superdiscount", 55}}, false},
		{"Multiple entries", fields{memdb.New()}, args{{10, "Superdiscount", 55}, {10, "Superdiscount2", 55}}, false},
		{"Duplicate entry", fields{memdb.New()}, args{{10, "Superdiscount", 55}, {10, "Superdiscount", 55}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}
			var anyErr error
			for _, coupon := range tt.args {
				if err := s.CreateCoupon(coupon.discount, coupon.code, coupon.minBasketValue); err != nil {
					anyErr = err
				}
			}
			if (anyErr != nil) != tt.wantErr {
				t.Errorf("CreateCoupon() error = %v, wantErr %v", anyErr, tt.wantErr)
			}
		})
	}
}

func TestService_GetCoupons(t *testing.T) {
	s := Service{
		repo: memdb.New(),
	}
	seed := []struct {
		discount       int
		code           string
		minBasketValue int
	}{
		{10, "Superdiscount", 55},
		{10, "Superdiscount2", 55},
	}
	for _, entry := range seed {
		err := s.CreateCoupon(entry.discount, entry.code, entry.minBasketValue)
		if err != nil {
			t.Fatal(err)
		}
	}
	type args struct {
		codes []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Get zero coupons", args{codes: []string{}}, false},
		{"Get single coupon", args{codes: []string{"Superdiscount"}}, false},
		{"Get multiple coupons", args{codes: []string{"Superdiscount", "Superdiscount2"}}, false},
		{"Get invalid coupon", args{codes: []string{"Superdiscount3"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetCoupons(tt.args.codes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoupons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
