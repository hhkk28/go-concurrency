## Exercise 6: Concurrent Web Crawler

### Description
Implement a web crawler that discovers and processes URLs concurrently. The crawler should:
- Start from a seed URL
- Discover links from pages (simulate)
- Process each page
- Respect concurrency limits
- Avoid revisiting URLs
- Handle timeouts

### Requirements
- Create a `Crawler` struct with configuration:
  - MaxConcurrentRequests (int)
  - MaxDepth (int)
  - Timeout (time.Duration)
- Implement crawling with these components:
  - URL frontier (to-visit queue)
  - Visited set (thread-safe)
  - Worker pool for concurrent fetching
  - Result collector
- Use context for cancellation
- Implement graceful shutdown
- Print crawl progress and statistics
- Handle errors gracefully (404, timeouts, etc.)

### Expected Output
```
Starting crawler...
  Max concurrent: 10
  Max depth: 3
  Timeout: 5s

Crawling: https://example.com (depth 0)
Crawling: https://example.com/about (depth 1)
Crawling: https://example.com/contact (depth 1)
Crawling: https://example.com/products (depth 1)
  Crawling: https://example.com/products/item1 (depth 2)
  Crawling: https://example.com/products/item2 (depth 2)
  Crawling: https://example.com/products/item3 (depth 2)
...

Crawl Statistics:
  Pages visited: 45
  URLs discovered: 67
  Errors: 3
  Total time: 2.5s
  Average time per page: 55ms
```

### Challenge
- Implement politeness policy (delay between requests to same domain)
- Add priority queue for URL frontier
