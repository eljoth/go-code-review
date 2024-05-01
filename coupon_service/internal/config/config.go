package config

import (
	"coupon_service/internal/api"
	"fmt"
	"log"

	"github.com/brumhard/alligotor"
)

type Config struct {
	API api.Config
}

func New() Config {
	cfg := Config{}
	if err := alligotor.Get(&cfg); err != nil {
		log.Fatal(err)
	}
	//missing port and host
	cfg.API.Port = 8080
	cfg.API.Host = "localhost"
	fmt.Println(cfg)
	return cfg
}
