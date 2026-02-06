## Exercise 4: Safe Concurrent Map with Real-time Statistics

### Description
Implement a thread-safe map that stores key-value pairs with real-time statistics tracking. The system should:
- Support concurrent reads and writes
- Track operation counts (reads, writes, deletes)
- Track the most popular keys
- Provide snapshot of current state
- Support atomic operations

### Requirements
- Create a `ConcurrentMap` type with methods:
  - `Set(key, value)` - thread-safe write
  - `Get(key)` - thread-safe read
  - `Delete(key)` - thread-safe delete
  - `GetStats()` - return operation statistics
  - `GetPopularKeys(n)` - return top N most accessed keys
- Use appropriate synchronization (Mutex or RWMutex)
- Track read/write/delete operations atomically
- Track key access frequency
- Implement graceful shutdown with `sync.Once`
- Create a statistics reporter that runs in background and prints stats every 5 seconds
- Simulate concurrent operations (10 writers, 50 readers)

### Expected Output
```
Concurrent operations running...
[5s] Stats: Reads: 245, Writes: 50, Deletes: 5, Keys: 45
[10s] Stats: Reads: 512, Writes: 98, Deletes: 12, Keys: 86
[15s] Stats: Reads: 768, Writes: 150, Deletes: 18, Keys: 132

Top 5 Popular Keys:
  1. "key42": accessed 45 times
  2. "key15": accessed 38 times
  3. "key7": accessed 32 times
  4. "key23": accessed 28 times
  5. "key89": accessed 25 times

Final Statistics:
  Total reads: 1024
  Total writes: 200
  Total deletes: 25
  Active keys: 175
```

### Challenge
- Implement cache eviction policy (LRU)
- Add TTL (time-to-live) for entries
