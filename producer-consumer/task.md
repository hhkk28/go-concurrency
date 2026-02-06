## Exercise 1: Producer-Consumer Pipeline

### Description
Implement a producer-consumer system where multiple producers generate data and multiple consumers process it. The pipeline should have three stages:
1. **Generator**: Produces a sequence of integers (0 to N)
2. **Squarer**: Receives integers and sends their squares
3. **Printer**: Receives squares and prints them

### Requirements
- Create 2 producer goroutines
- Create 3 consumer goroutines
- Use channels to communicate between stages
- Ensure all goroutines clean up properly
- No goroutine should leak
- The pipeline should be fully drained (all values processed)
- Handle shutdown gracefully when all producers finish

### Expected Output
```
Received square: 0
Received square: 1
Received square: 4
...
```

### Challenge
- Add a timeout context that shuts down the entire pipeline if processing takes too long