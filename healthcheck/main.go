package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

const defaultProbeTimeout = 2 * time.Second

func main() {
	os.Exit(runProbe())
}

func runProbe() int {
	probeURL := os.Getenv("PROBE_URL")
	if probeURL == "" {
		fmt.Fprintln(os.Stderr, "probe failed: PROBE_URL is not set")
		return 1
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultProbeTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, probeURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "probe request error: %v\n", err)
		return 1
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "probe failed: %v\n", err)
		return 1
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "probe failed with status %d\n", resp.StatusCode)
		return 1
	}

	return 0
}
