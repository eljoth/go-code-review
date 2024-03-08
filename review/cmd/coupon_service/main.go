package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

func main() {
	svc := service.New(repo)
	router := api.New(cfg.Api, svc)
	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Println("Starting Coupon service server")
		router.Start()
	}()
	<-term
	log.Println("Closing Coupon service server")
	router.Close()
}
