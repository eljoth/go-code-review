package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("failed to load config", slog.Any("err", err))
		os.Exit(1)
	}

	logLevel := slog.LevelDebug
	if cfg.Runtime.LogLevel != "" {
		logLevel, err = parseLogLevel(cfg.Runtime.LogLevel)
		if err != nil {
			slog.Error("failed to parse log level", slog.Any("err", err))
			os.Exit(1)
		}
	}
	slog.SetLogLoggerLevel(logLevel)

	svc := service.New(memdb.New())
	api := api.New(cfg.API, svc)

	go func() {
		slog.Info("starting coupon service server", slog.String("host", cfg.API.Host), slog.Int("port", cfg.API.Port))
		if err := api.Start(); err != nil {
			slog.Error("failed to start coupon service server", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	slog.Info("shutting down coupon service server")
	err = api.Close()
	if err != nil {
		slog.Error("failed to shutdown coupon service server gracefully", slog.Any("err", err))
	}

	slog.Info("bye!")
}

func parseLogLevel(s string) (slog.Level, error) {
	var level slog.Level
	var err = level.UnmarshalText([]byte(s))
	return level, err
}
