package main

import (
	"fmt"
	"sort"
	"sync"
)

func generator(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range numbers {
			out <- v
		}
	}()
	return out
}

func worker(id int, input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range input {
			fmt.Println("Worker", id, "processed", v, ", result:", v*v)
			out <- v * v
		}
		fmt.Println("End of inputs in Worker:", id)
	}()
	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, c := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	source := generator(numbers)
	workers := make([]<-chan int, 3)
	for i := range 3 {
		workers[i] = worker(i+1, source)
	}
	merged := fanIn(workers...)
	var results []int
	for v := range merged {
		results = append(results, v)
	}
	sort.Ints(results)
	for _, v := range results {
		fmt.Println("Result:", v)
	}
}
