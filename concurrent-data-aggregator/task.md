## Exercise 3: Concurrent Data Aggregator

### Description
Create a system that concurrently fetches data from multiple sources and aggregates results. The system should:
- Fetch data from 5 different "sources" (simulate with functions)
- Some sources may be slow, some fast
- Some may fail
- Aggregate all successful results
- Timeout after a maximum wait time
- Return partial results if some sources fail

### Requirements
- Implement a `Source` struct with:
  - Name (string)
  - Latency (time.Duration - simulate network delay)
  - ShouldFail (bool - simulate failures)
- Fetch from all 5 sources concurrently
- Use `select` with a timeout
- Collect results from all successful fetches
- Use context for cancellation
- Implement retry logic for failed sources (max 3 retries)
- Return aggregated results with metadata

### Expected Output
```
Fetching from Source A...
Fetching from Source B...
Fetching from Source C...
Fetching from Source D...
Fetching from Source E...

Source A: data-a (took 100ms)
Source B: FAILED (retry 1/3)
Source B: data-b (took 200ms, retry success)
Source C: data-c (took 50ms)
Source D: TIMEOUT after 300ms
Source E: data-e (took 150ms)

Aggregation Summary:
  Successful: 4/5
  Failed: 1/5
  Total time: 300ms
  Data: [data-a, data-b, data-c, data-e]
```

### Challenge
- Implement exponential backoff for retries
- Add circuit breaker pattern (stop retrying after N consecutive failures)
