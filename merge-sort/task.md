## Bonus Challenge: Concurrent Merge Sort

### Description
Implement a merge sort algorithm that uses goroutines to sort sub-arrays concurrently.

### Requirements
- Implement concurrent merge sort
- Split array into chunks
- Sort chunks concurrently
- Merge sorted chunks
- Limit maximum concurrent goroutines
- Benchmark against sequential merge sort

### Expected
```
Sequential sort: 450ms
Concurrent sort (4 workers): 180ms
Concurrent sort (8 workers): 120ms
Speedup: 3.75x
```
