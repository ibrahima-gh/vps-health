package main

import (
	"fmt"
	"time"

	"github.com/ibrahima-gh/vps-health/internal/checker"
	"github.com/ibrahima-gh/vps-health/internal/config"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err)
		return
	}

	if cfg == nil {
		fmt.Println("config is nil (Load not yet implemented)")
		return
	}

	timeout := time.Duration(cfg.TimeoutSeconds) * time.Second
	results := checker.CheckAll(cfg.Targets, timeout)

	_ = results
}
