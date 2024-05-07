package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"fmt"

	// "runtime"
	"time"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

// func init() {
// 	if runtime.NumCPU() != 32 {
// 		panic("this api is meant to be run on 32 core machines")
// 	}
// }

func main() {
	svc := service.New(repo)
	newAPI := api.New(cfg.API, svc)
	defer newAPI.Close()

	fmt.Println("Starting Coupon service server")
	newAPI.Start()

	<-time.After(1 * time.Hour * 24 * 365)
	fmt.Println("Coupon service server alive for a year, closing")
	newAPI.Close()
}
