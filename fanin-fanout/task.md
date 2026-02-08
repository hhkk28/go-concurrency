# Fan-in and Fan-out Exercise

## Objective
Practice implementing fan-out (splitting work across multiple goroutines) and fan-in (merging results from multiple channels) patterns in Go.

## Problem Description

You are building a number processing system that:
1. **Fan-out**: Distribute numbers from a source across multiple worker goroutines
2. **Process**: Each worker squares the numbers it receives
3. **Fan-in**: Merge all results from workers into a single channel

## Task

Complete the implementation in `main.go` with the following functions:

### 1. `generator(numbers []int) <-chan int`
- Takes a slice of integers
- Sends each number to a channel
- Returns a read-only channel
- Closes the channel when all numbers are sent

### 2. `worker(id int, input <-chan int) <-chan int`
- A worker goroutine that:
  - Receives numbers from input channel
  - Squares each number
  - Sends squared numbers to output channel
  - Logs its activity: "Worker X processed N, result: R"
  - Closes output channel when done

### 3. `fanIn(channels ...<-chan int) <-chan int`
- Takes multiple channels as variadic argument
- Creates a single output channel
- Uses goroutines to merge all inputs into one output
- Closes output channel when all inputs are exhausted

## Requirements

- Use only one worker per input channel in `fanIn`
- All channels must be properly closed
- Handle any number of workers (2, 3, 4, etc.)
- Print worker activities to show concurrent execution

## Example Usage

```go
numbers := []int{1, 2, 3, 4, 5, 6}
source := generator(numbers)

// Fan-out: create 3 workers
worker1 := worker(1, source)
worker2 := worker(2, source)
worker3 := worker(3, source)

// Fan-in: merge all results
results := fanIn(worker1, worker2, worker3)

// Print results
for result := range results {
    fmt.Println(result)
}
```

## Expected Output (order may vary due to concurrency)

```
Worker 1 processed 1, result: 1
Worker 2 processed 2, result: 4
Worker 3 processed 3, result: 9
Worker 1 processed 4, result: 16
Worker 2 processed 5, result: 25
Worker 3 processed 6, result: 36
1
4
9
16
25
36
```

## Hints

1. **Generator**: Simple loop, defer close
2. **Worker**: range over input, close output when done
3. **Fan-in**: Use `select` statement with multiple cases, or launch a goroutine for each input channel

## Learning Objectives

- Understand how to distribute work (fan-out)
- Learn to merge results from multiple goroutines (fan-in)
- Practice channel synchronization
- Use `select` for multiplexing channels

## Extension

After completing the basic version:
- Add a context to cancel workers mid-execution
- Add metrics (how many numbers each worker processed)
- Implement worker pools with dynamic sizing
