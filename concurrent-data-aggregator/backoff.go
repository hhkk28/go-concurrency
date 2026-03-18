package main

import (
	"context"
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

var counter int32

func RetryWithBackoff(ctx context.Context, fn func() error, onRetry func(attempt int32), onFailure func(err error)) {
	for {
		if ctx.Err() != nil {
			onFailure(fmt.Errorf("Context cancelled"))
			return
		}
		if err := fn(); err == nil {
			return
		} else if counter == MAX_RETRIES {
			onFailure(fmt.Errorf("Maximum retries exceeded"))
			return
		}
		onRetry(counter)
		atomic.AddInt32(&counter, 1)
		delay := min(time.Duration(math.Pow(2, float64(counter))*float64(time.Second)), time.Duration(MAX_BACKOFF))
		select {
		case <-ctx.Done():
			onFailure(fmt.Errorf("Entire function timed out"))
			return
		case <-time.After(delay):
		}
	}
}
