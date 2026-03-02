package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	maxRetries    = 3
	maxBackoff    = 30 * time.Second
	globalTimeout = 5 * time.Second
)

type Source struct {
	Name       string
	Latency    time.Duration
	ShouldFail bool
}

type Result struct {
	Source   string
	Err      error
	Data     string
	Duration time.Duration
}

var SourceList = []Source{
	{"AlphaStream", 750 * time.Millisecond, false},
	{"BetaNode", 2100 * time.Millisecond, false},
	{"GammaFeed", 500 * time.Millisecond, true},
	{"DeltaPulse", 3000 * time.Millisecond, false},
	{"EpsilonCache", 1200 * time.Millisecond, false},
}

func DisplaySummary(out <-chan Result) {
	successful := 0
	failed := 0
	var totalTime time.Duration
	total := 0
	result := []string{}
	for res := range out {
		if err := res.Err; err == nil {
			result = append(result, res.Data)
			successful++
		} else {
			failed++
		}
		totalTime += res.Duration
		total++
	}
	fmt.Println("Aggregation Summary:")
	fmt.Printf("\tSuccessful: %d/%d\n", successful, total)
	fmt.Printf("\tFailed: %d/%d\n", failed, total)
	fmt.Printf("\tTotal time: %v\n", totalTime)
	fmt.Printf("\tData:%v\n", result)
}

func Fetch(ctx context.Context, source Source, out chan<- Result) error {
	fmt.Printf("Fetching from %s\n", source.Name)
	start := time.Now()
	shouldSucceed := rand.Intn(maxRetries)%2 == 0
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

func RetryWithBackoff(ctx context.Context, fn func() error, onRetry func(attempt int), onFailure func(err error)) {
	for attempted := range maxRetries {
		if ctx.Err() != nil {
			onFailure(fmt.Errorf("Context cancelled"))
			return
		}
		if err := fn(); err == nil {
			return
		} else if attempted == maxRetries-1 {
			onFailure(fmt.Errorf("Maximum retries exceeded"))
			return
		}
		onRetry(attempted)
		delay := min(time.Duration(math.Pow(2, float64(attempted))*float64(time.Second)), time.Duration(maxBackoff))
		select {
		case <-ctx.Done():
			onFailure(fmt.Errorf("Entire function timed out"))
			return
		case <-time.After(delay):
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), globalTimeout)
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
			}, func(count int) {
				fmt.Printf("Source %s FAILED (retry %d/%d)\n", source.Name, count+1, maxRetries)
			}, func(err error) {
				out <- Result{Source: source.Name, Err: err, Duration: time.Since(start), Data: ""}
			})
		}()
	}
	wg.Wait()
	close(out)
	DisplaySummary(out)
}
