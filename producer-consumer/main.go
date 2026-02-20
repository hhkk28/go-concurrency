package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var N = 50
var NUM_OF_PRODUCERS = 4
var NUM_OF_CONSUMERS = 6
var TIMEOUT = 5 * time.Second
var COUNTER atomic.Int64

func Generator(id int, intChannel chan<- int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// To prevent from adding multiple 0 to the channel, we are incrementing right
		// away and subtracting 1 from it.
		value := int(COUNTER.Add(1)) - 1
		if value >= N {
			return
		}
		select {
		case <-ctx.Done():
			return
		case intChannel <- value:
		}
	}
}

func Squarer(intChannel <-chan int, squaredChannel chan<- int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-intChannel:
			if ok {
				select {
				case <-ctx.Done():
					return
				case squaredChannel <- v * v:
				}
			} else {
				return
			}
		}
	}
}

func Printer(squaredChannel <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-squaredChannel:
			if ok {
				fmt.Println("Received square:", v)
			} else {
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()
	intChannel := make(chan int)
	squaredChannel := make(chan int)
	var genWg sync.WaitGroup
	for i := range NUM_OF_PRODUCERS {
		genWg.Add(1)
		go Generator(i, intChannel, ctx, &genWg)
	}
	go func() {
		genWg.Wait()
		close(intChannel)
	}()

	var sqWg sync.WaitGroup
	for range NUM_OF_CONSUMERS {
		sqWg.Add(1)
		go Squarer(intChannel, squaredChannel, ctx, &sqWg)
	}
	go func() {
		sqWg.Wait()
		close(squaredChannel)
	}()

	var printWg sync.WaitGroup
	for range NUM_OF_CONSUMERS {
		printWg.Add(1)
		go Printer(squaredChannel, ctx, &printWg)
	}
	printWg.Wait()
	fmt.Println("All done")
}
