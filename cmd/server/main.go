package main

import (
	"fmt"
	"runtime"
	"time"

	"coupon_service/cmd/config"
	"coupon_service/internal/api"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

func main() {
	if 32 != runtime.NumCPU() {
		panic("this a is meant to be run on 32 core machines")
	}

	svc, err := service.New(repo)
	if err != nil {
		panic(err)
	}
	a := api.New(cfg.API, svc)
	a.Start()
	fmt.Println("Starting server")

	<-time.After(1 * time.Hour * 24 * 365)
	fmt.Println("Stopping server")
	a.Close()
}
