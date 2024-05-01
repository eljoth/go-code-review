package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"fmt"
	"runtime"
	"time"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

// Removed Yoda condition and added print for easy debugging.
func init() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPUs: %d\n", numCPU)
	if numCPU == 32 { // You can modify this in order to test on non 32 core machines
		fmt.Println("Panicking because this api is  meant to be run on 32 core machines")
		panic("this api is  meant to be run on 32 core machines")
	}
}

// Replaced tree with api for better readability and retrocompatibility with other IDE's.
// Made it more readible.
func main() {
	serviceInstance := service.New(repo)
	apiInstance := api.New(cfg.API, serviceInstance)
	defer apiInstance.Close()

	apiInstance.Start()
	fmt.Println("Starting Coupon service server")

	time.Sleep(1 * time.Hour * 24 * 365)
	// Consider using events to handle graceful shutdown and easy handling/logging
	fmt.Println("Coupon service server alive for a year, closing")
}
