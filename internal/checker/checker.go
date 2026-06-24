package checker

import (
	"time"

	"github.com/ibrahima-gh/vps-health/internal/config"
)

type Result struct {
	Name       string
	URL        string
	StatusCode int
	Latency    time.Duration
	SSLExpiry  time.Time
	Err        error
}

func CheckAll(targets []config.Target, timeout time.Duration) []Result {
	return nil
}

func check(target config.Target, timeout time.Duration) Result {
	return Result{}
}
