package config

import (
	"coupon_service/internal/api"
	"fmt"

	"log/slog"

	"github.com/brumhard/alligotor"
)

type Config struct {
	Runtime Runtime
	API     api.Config
}

type Runtime struct {
	LogLevel string
}

func New() (*Config, error) {
	cfg := Config{
		Runtime: Runtime{
			LogLevel: slog.LevelDebug.String(),
		},
		API: api.Config{
			Host: "localhost",
			Port: 8080,
		},
	}
	if err := alligotor.Get(&cfg); err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	return &cfg, nil
}
