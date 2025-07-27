package logger

import (
	"TaskTrackerPlus/config"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func New(cfg config.Config) *slog.Logger {
	const op = "pkg.logger.New"
	switch cfg.Env {
	case envLocal:
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		return logger
	case envProd:
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		return logger
	default:
		log.Fatalf("%s:%s", op, "env variable doesn't fit any")
	}
	return nil
}
