package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

func RetryWithBackoff(ctx context.Context, fn func() error, onRetry func(attempt int), onFailure func(err error)) {
	for attempted := range MAX_RETRIES {
		if ctx.Err() != nil {
			onFailure(fmt.Errorf("Context cancelled"))
			return
		}
		if err := fn(); err == nil {
			return
		} else if attempted == MAX_RETRIES-1 {
			onFailure(fmt.Errorf("Maximum retries exceeded"))
			return
		}
		onRetry(attempted)
		delay := min(time.Duration(math.Pow(2, float64(attempted))*float64(time.Second)), time.Duration(MAX_BACKOFF))
		select {
		case <-ctx.Done():
			onFailure(fmt.Errorf("Entire function timed out"))
			return
		case <-time.After(delay):
		}
	}
}
