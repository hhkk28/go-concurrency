package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Fetch(ctx context.Context, source Source, out chan<- Result) error {
	fmt.Printf("Fetching from %s\n", source.Name)
	start := time.Now()
	shouldSucceed := rand.Intn(MAX_RETRIES)%2 == 0
	if shouldSucceed && !source.ShouldFail {
		select {
		case <-ctx.Done():
			return fmt.Errorf("Entire function timed out")
		case <-time.After(source.Latency):
			fmt.Printf("Source %s: Data from %s (took %v)\n", source.Name, source.Name, time.Since(start))
			out <- Result{Source: source.Name, Data: fmt.Sprintf("Data from %s", source.Name), Err: nil, Duration: time.Since(start)}
			return nil
		}
	}
	return fmt.Errorf("Failed to fetch")
}
