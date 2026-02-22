package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type StatusState int

const (
	StatusCompleted StatusState = iota
	StatusFailed
	StatusTimedOut
)

const NUM_OF_WORKERS = 5

func arrayToChannel(arr []Job) chan Job {
	ch := make(chan Job, len(arr))
	go func() {
		defer close(ch)
		for _, e := range arr {
			ch <- e
		}
	}()
	return ch
}

func processJobs(jobsChannel chan Job, resultsChannel chan StatusState, ctx context.Context, wg *sync.WaitGroup) {
	workerChannel := make(chan struct{}, NUM_OF_WORKERS)
	for j := range jobsChannel {
		go func() {
			defer wg.Done()
			select {
			case workerChannel <- struct{}{}:
				defer func() {
					<-workerChannel
				}()
				time.Sleep(j.Duration)
				completed := rand.Intn(5) >= 1
				if completed {
					fmt.Println(j.Payload, "completed successfully")
					resultsChannel <- StatusCompleted
				} else {
					fmt.Println(j.Payload, "failed: unexpected error")
					resultsChannel <- StatusFailed
				}
			case <-ctx.Done():
				fmt.Println(j.Payload, "stopped because of timeout")
				resultsChannel <- StatusTimedOut
				return
			}
		}()
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jobsChannel := arrayToChannel(Jobs)
	resultsChannel := make(chan StatusState, len(Jobs))
	currentTime := time.Now()
	var duration time.Duration = 0
	wg.Add(len(Jobs))
	go func() {
		processJobs(jobsChannel, resultsChannel, ctx, &wg)
	}()
	wg.Wait()
	duration = time.Since(currentTime)
	close(resultsChannel)
	completed, total, failed, timedOut := 0, 0, 0, 0
	for v := range resultsChannel {
		total += 1
		switch v {
		case StatusCompleted:
			completed += 1
		case StatusFailed:
			failed += 1
		case StatusTimedOut:
			timedOut += 1
		}
	}
	fmt.Println("Execution status:")
	fmt.Println("Total:\t", total)
	fmt.Println("Completed:\t", completed)
	fmt.Println("Failed:\t", failed)
	fmt.Println("Timed out:\t", timedOut)
	fmt.Println("Duration:\t", duration)
}
