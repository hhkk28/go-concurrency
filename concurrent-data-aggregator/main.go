package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), GLOBAL_TIMEOUT)
	defer cancel()
	var wg sync.WaitGroup
	out := make(chan Result, len(SourceList))
	wg.Add(len(SourceList))
	for _, source := range SourceList {
		start := time.Now()
		go func() {
			defer wg.Done()
			RetryWithBackoff(ctx, func() error {
				return Fetch(ctx, source, out)
			}, func(count int32) {
				fmt.Printf("Source %s FAILED (retry %d/%d)\n", source.Name, count+1, MAX_RETRIES)
			}, func(err error) {
				out <- Result{Source: source.Name, Err: err, Duration: time.Since(start), Data: ""}
			})
		}()
	}
	wg.Wait()
	close(out)
	DisplaySummary(out)
}
