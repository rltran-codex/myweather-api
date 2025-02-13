package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestRateLimit(t *testing.T) {
	url := "http://localhost:8080/api/v1/weather"

	// Define how many requests should exceed the limit (based on rate limiter settings)
	totalRequests := 10
	successCount := 0
	failedCount := 0

	for i := 0; i < totalRequests; i++ {
		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		// Read response body (optional, but helps avoid connection reuse issues)
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		// Log response
		fmt.Printf("Request %d: Status %d, Body: %s\n", i+1, resp.StatusCode, string(body))

		// Check if rate limit is hit
		if resp.StatusCode == http.StatusTooManyRequests {
			failedCount++
		} else if resp.StatusCode == http.StatusOK {
			successCount++
		}

		// Small delay between requests (optional)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("Total successful requests: %d, Total rate limited: %d\n", successCount, failedCount)

	// Assert that some requests were rate-limited
	if failedCount == 0 {
		t.Errorf("Expected some requests to be rate-limited, but none were.")
	}
}
