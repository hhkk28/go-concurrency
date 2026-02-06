# Go Concurrency Mastery Exercises

A comprehensive collection of hands-on exercises to master Go's concurrency primitives and patterns. Each exercise demonstrates real-world use cases and best practices for writing concurrent Go programs.

## 📚 Overview

This repository contains 7 core exercises plus 1 bonus challenge, each focusing on different aspects of Go concurrency:

| Exercise | Focus Area | Difficulty |
|----------|-----------|------------|
| [Producer-Consumer Pipeline](./producer-consumer/) | Channels, goroutines, pipeline patterns | Beginner |
| [Rate-Limited Worker Pool](./rate-limited-worker-pool/) | Worker pools, semaphores, rate limiting | Beginner |
| [Concurrent Data Aggregator](./concurrent-data-aggregator/) | Context, timeouts, error handling | Intermediate |
| [Safe Concurrent Map](./map-with-stats/) | Mutex, RWMutex, atomic operations | Intermediate |
| [Pub/Sub Event System](./pub-sub-system/) | Fan-out/fan-in, message routing | Intermediate |
| [Concurrent Web Crawler](./web-crawler/) | Coordination, state management | Advanced |
| [Orchestration with Context](./orchestration/) | Dependency graphs, error propagation | Advanced |
| [Concurrent Merge Sort](./merge-sort/) (Bonus) | Divide-and-conquer, benchmarking | Advanced |

## 🎯 Prerequisites

- Understanding of goroutines
- Knowledge of channels (buffered and unbuffered)
- Familiarity with `sync` package (WaitGroup, Mutex, RWMutex, Once, Atomic)
- Understanding of `select` statements
- Knowledge of context package for cancellation
- Understanding of worker pool patterns

## 🚀 Getting Started

1. **Clone or navigate to the exercise directory**
   ```bash
   cd exercise-name/
   ```

2. **Read the task description**
   ```bash
   cat task.md
   ```

3. **Implement your solution**
   - Create `main.go` with your implementation
   - Include proper error handling and resource cleanup

4. **Write tests**
   - Create `main_test.go` with comprehensive tests
   - Test edge cases and concurrent scenarios

5. **Run tests**
   ```bash
   go test -v
   go test -race  # Check for race conditions
   ```

## 📖 Exercise Details

### Exercise 1: Producer-Consumer Pipeline
Implement a multi-stage pipeline where producers generate data and consumers process it through channels. Learn proper goroutine lifecycle management and pipeline patterns.

### Exercise 2: Rate-Limited Worker Pool
Build a worker pool with rate limiting that processes jobs concurrently while respecting resource limits. Covers semaphores, job queuing, and statistics tracking.

### Exercise 3: Concurrent Data Aggregator
Create a system that fetches data from multiple sources concurrently, handles timeouts and failures, and aggregates results. Includes retry logic and exponential backoff.

### Exercise 4: Safe Concurrent Map with Real-time Statistics
Implement a thread-safe map with operation statistics tracking. Covers RWMutex for read-heavy workloads, atomic counters, and background reporting.

### Exercise 5: Pub/Sub Event System
Build a publish-subscribe system with topic filtering, wildcard matching, and backpressure handling. Demonstrates fan-out/fan-in patterns and non-blocking message delivery.

### Exercise 6: Concurrent Web Crawler
Implement a concurrent web crawler with URL frontier management, visited set tracking, and graceful shutdown. Covers coordination patterns and error handling.

### Exercise 7: Orchestration with Context and Error Propagation
Create a task orchestration system that executes tasks respecting dependency graphs, propagates errors, and supports cancellation. Demonstrates complex coordination patterns.

### Bonus: Concurrent Merge Sort
Implement concurrent merge sort with configurable worker counts and benchmarking against sequential implementation.

## ✅ Testing Your Solutions

For each exercise, ensure you:

1. **Write comprehensive unit tests**
   - Test core functionality
   - Test edge cases (empty inputs, nil values, timeouts, failures)
   - Test concurrent scenarios

2. **Check for race conditions**
   ```bash
   go test -race
   ```

3. **Verify no goroutine leaks**
   - Check goroutine counts before/after
   - Ensure proper cleanup on shutdown

4. **Measure performance**
   - Benchmark where applicable
   - Compare with sequential implementations

## 📊 Evaluation Criteria

Your solutions will be evaluated on:

- ✅ **Correctness** - does it work as specified?
- ✅ **Concurrency Safety** - no race conditions or deadlocks
- ✅ **Resource Management** - no goroutine leaks, proper cleanup
- ✅ **Error Handling** - graceful degradation, proper error propagation
- ✅ **Code Quality** - idiomatic Go, proper naming, documentation
- ✅ **Performance** - efficient use of concurrency
- ✅ **Test Coverage** - comprehensive test cases

## 💡 Tips for Success

1. **Start sequential, then parallel** - Begin with a working sequential solution, then add concurrency
2. **Always test with race detector** - Use `go test -race` to catch data races early
3. **Use defer for cleanup** - Ensure channels are closed and WaitGroups are properly decremented
4. **Think about edge cases** - What happens if channels are closed? What if context is cancelled?
5. **Use meaningful names** - Clear variable and function names make concurrent code easier to understand
6. **Add comments** - Document synchronization patterns and invariants
7. **Test with different configurations** - Try different worker counts, timeouts, and data sizes
8. **Use profiling tools** - `pprof` can help identify performance bottlenecks

## 🛠️ Common Patterns

### Worker Pool
```go
func worker(id int, jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        results <- process(job)
    }
}
```

### Fan-out/Fan-in
```go
// Distribute work to multiple goroutines
for _, w := range workers {
    go w(input)
}

// Collect results
for i := 0; i < len(results); i++ {
    <-results
}
```

### Context Cancellation
```go
ctx, cancel := context.WithTimeout(context.Background(), timeout)
defer cancel()

select {
case result := <-ch:
    return result
case <-ctx.Done():
    return ctx.Err()
}
```

## 📦 Recommended Tools

- **Race Detector**: `go test -race`
- **Benchmarking**: `go test -bench=.`
- **Profiling**: `pprof`, `trace`
- **Linting**: `golangci-lint`
- **Static Analysis**: `go vet`

## 📚 Resources & Documentation

### Official Go Documentation
- [Go Concurrency Patterns](https://go.dev/doc/codelab/concurrency) - Official codelab on concurrency patterns
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency) - Best practices for concurrent code
- [The Go Memory Model](https://go.dev/ref/mem) - Understanding memory synchronization in Go
- [Package: sync](https://pkg.go.dev/sync) - Documentation for sync package (Mutex, WaitGroup, etc.)
- [Package: context](https://pkg.go.dev/context) - Context package for cancellation and deadlines

### Go by Example
- [Goroutines](https://gobyexample.com/goroutines) - Basic goroutine usage
- [Channels](https://gobyexample.com/channels) - Channel basics
- [Buffered Channels](https://gobyexample.com/buffered-channels) - Buffered channel patterns
- [Channel Directions](https://gobyexample.com/channel-directions) - Type-safe channel usage
- [Channel Select](https://gobyexample.com/channel-select) - Non-blocking channel operations
- [Timeouts](https://gobyexample.com/timeouts) - Context timeouts and cancellation
- [Worker Pools](https://gobyexample.com/worker-pools) - Worker pool pattern
- [Rate Limiting](https://gobyexample.com/rate-limiting) - Rate limiting patterns

### Go Blog Posts
- [Share Memory By Communicating](https://go.dev/blog/codelab-share) - Philosophy behind Go's concurrency model
- [Go Concurrency Patterns: Context](https://go.dev/blog/context) - Context package usage
- [Go Concurrency Patterns: Timing out, moving on](https://go.dev/blog/pipelines) - Pipeline patterns and timeouts

**Happy Coding! 🚀**

Remember: Concurrency is not parallelism, but it's about managing and composing independent processes. Master Go's concurrency primitives and you'll unlock the full potential of the language.