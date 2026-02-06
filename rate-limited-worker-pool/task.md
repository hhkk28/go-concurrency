## Exercise 2: Rate-Limited Worker Pool

### Description
Implement a worker pool that processes jobs with rate limiting. The system should:
- Accept jobs from a channel
- Limit concurrent jobs to a maximum of 5 workers
- Process each job (simulate with sleep)
- Return results to a results channel
- Track and report statistics (completed, failed, total jobs)

### Requirements
- Implement a `Job` struct with:
  - ID (int)
  - Payload (string)
  - Duration (time.Duration - simulate work time)
- Implement rate limiting using either:
  - A buffered channel as semaphore
  - Or ticker-based throttling
- Handle job failures gracefully (some jobs should randomly fail)
- Return results through a channel
- Print final statistics when all jobs are complete
- Support cancellation via context

### Example Job Data
```go
jobs := []Job{
    {ID: 1, Payload: "Process image", Duration: 100 * time.Millisecond},
    {ID: 2, Payload: "Send email", Duration: 50 * time.Millisecond},
    // ... more jobs
}
```

### Expected Output
```
Job 1 completed successfully
Job 2 failed: simulated error
Job 3 completed successfully
...
Statistics:
  Total jobs: 20
  Completed: 18
  Failed: 2
  Duration: 1.2s
```

### Challenge
- Add a dynamic worker count that can be adjusted at runtime
- Implement priority queue for jobs
