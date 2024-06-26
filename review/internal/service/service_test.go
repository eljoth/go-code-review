package service

import (
	"coupon_service/internal/repository"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service/entity"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		repo repository.Repository
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
		repo    repository.Repository
		coupons []entity.Coupon
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
		{
			"Apply 10%",
			fields{memdb.New(), []entity.Coupon{{ID: "1", Code: "10_off", Discount: 10, MinBasketValue: 3}}},
			args{entity.Basket{Value: 100}, "10_off"},
			&entity.Basket{Value: 100, AppliedDiscount: 10, ApplicationSuccessful: true},
			false,
		},
		{
			"Apply 50%",
			fields{memdb.New(), []entity.Coupon{{ID: "1", Code: "50_off", Discount: 50, MinBasketValue: 3}}},
			args{entity.Basket{Value: 100}, "50_off"},
			&entity.Basket{Value: 100, AppliedDiscount: 50, ApplicationSuccessful: true},
			false,
		},
		{
			"Fail to apply 10%",
			fields{memdb.New(), []entity.Coupon{{ID: "1", Code: "10_off", Discount: 10, MinBasketValue: 200}}},
			args{entity.Basket{Value: 100}, "10_off"},
			&entity.Basket{Value: 100, AppliedDiscount: 0, ApplicationSuccessful: false},
			false,
		},
		{
			"Unexisting coupon",
			fields{memdb.New(), nil},
			args{entity.Basket{Value: 100}, "10_off"},
			nil,
			true,
		},
		{
			"Negative value basket",
			fields{memdb.New(), []entity.Coupon{{ID: "1", Code: "10_off", Discount: 10, MinBasketValue: 0}}},
			args{entity.Basket{Value: -10}, "10_off"},
			nil,
			true,
		},
		{
			"Zero value basket",
			fields{memdb.New(), []entity.Coupon{{ID: "1", Code: "10_off", Discount: 10, MinBasketValue: 0}}},
			args{entity.Basket{Value: 0}, "10_off"},
			&entity.Basket{Value: 0, AppliedDiscount: 0, ApplicationSuccessful: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}
			for _, c := range tt.fields.coupons {
				err := s.repo.Save(c)
				if err != nil {
					t.Errorf("failed to save coupon: %v", c)
				}
			}

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
		repo            repository.Repository
		existingCoupons []entity.Coupon
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
		wantC   *entity.Coupon
		wantErr bool
	}{
		{
			"Create 10%",
			fields{memdb.New(), nil},
			args{10, "Superdiscount", 55},
			&entity.Coupon{Code: "Superdiscount", Discount: 10, MinBasketValue: 55},
			false,
		},
		{
			"Create 50%",
			fields{memdb.New(), []entity.Coupon{{ID: "1", Code: "Superdiscount", Discount: 10, MinBasketValue: 200}}},
			args{50, "Superdiscount50", 100},
			&entity.Coupon{Code: "Superdiscount50", Discount: 50, MinBasketValue: 100},
			false,
		},
		{
			"Already existing error",
			fields{memdb.New(), []entity.Coupon{{ID: "1", Code: "Superdiscount", Discount: 10, MinBasketValue: 200}}},
			args{10, "Superdiscount", 55},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}
			for _, c := range tt.fields.existingCoupons {
				err := s.repo.Save(c)
				if err != nil {
					t.Errorf("failed to save coupon: %v", c)
					return
				}
			}

			gotC, err := s.CreateCoupon(tt.args.discount, tt.args.code, tt.args.minBasketValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				if gotC.ID == "" {
					t.Errorf("CreateCoupon() returned coupon with empty ID")
				}
				gotC.ID = ""
			}

			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("ApplyCoupon() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
