package main

import (
	"TaskTrackerPlus/config"
)

func main() {
	cfg := config.MustLoad()
	_ = cfg
}
