package checker

import (
	"crypto/tls"
	"fmt"
	"net/http"
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

func CheckAll(targets []config.Target, timeout time.Duration, onResult func(Result)) {
	for _, t := range targets {
		onResult(check(t, timeout))
	}
}

func check(target config.Target, timeout time.Duration) Result {
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
	}

	start := time.Now()
	resp, err := client.Get(target.URL)
	latency := time.Since(start)

	if err != nil {
		return Result{Name: target.Name, URL: target.URL, Latency: latency, Err: fmt.Errorf("request failed: %w", err)}
	}
	defer resp.Body.Close()

	r := Result{
		Name:       target.Name,
		URL:        target.URL,
		StatusCode: resp.StatusCode,
		Latency:    latency,
	}

	if resp.TLS != nil && len(resp.TLS.PeerCertificates) > 0 {
		r.SSLExpiry = resp.TLS.PeerCertificates[0].NotAfter
	}

	return r
}
