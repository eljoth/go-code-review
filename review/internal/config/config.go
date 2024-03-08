package config

import (
	"coupon_service/internal/api"
	"log"

	"github.com/brumhard/alligotor"
)

type Config struct {
	Api api.Config
}

func New() Config {
	cfg := Config{
		Api: api.Config{
			Host: "127.0.0.1",
			Port: 8080,
		},
	}
	if err := alligotor.Get(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
