package main

import (
	"TaskTrackerPlus/config"
	"TaskTrackerPlus/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.New(cfg)
	logger.Error("Test")
	_ = cfg
}
