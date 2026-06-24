package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/ibrahima-gh/vps-health/internal/checker"
	"github.com/ibrahima-gh/vps-health/internal/config"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		color.Red("error: %v", err)
		return
	}

	fmt.Printf("checking %d targets (timeout: %ds)...\n\n", len(cfg.Targets), cfg.TimeoutSeconds)

	timeout := time.Duration(cfg.TimeoutSeconds) * time.Second
	results := checker.CheckAll(cfg.Targets, timeout)

	green := color.New(color.FgGreen).SprintfFunc()
	red := color.New(color.FgRed).SprintfFunc()
	yellow := color.New(color.FgYellow).SprintfFunc()

	for _, r := range results {
		if r.Err != nil {
			fmt.Printf("%s  %-20s  %s\n", red("✗"), r.Name, red(r.Err.Error()))
			continue
		}

		ssl := "no TLS"
		if !r.SSLExpiry.IsZero() {
			daysLeft := int(time.Until(r.SSLExpiry).Hours() / 24)
			if daysLeft < 14 {
				ssl = yellow("SSL expires in %dd", daysLeft)
			} else {
				ssl = green("SSL ok (%dd)", daysLeft)
			}
		}

		fmt.Printf("%s  %-20s  %s  %s  %s\n",
			green("✓"),
			r.Name,
			green("%d", r.StatusCode),
			r.Latency.Round(time.Millisecond),
			ssl,
		)
	}
}
