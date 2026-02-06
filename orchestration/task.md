
## Exercise 7: Orchestration with Context and Error Propagation

### Description
Implement a task orchestration system where multiple tasks run in a dependency graph. Tasks can:
- Run in parallel when independent
- Run sequentially when dependent
- Fail and propagate errors to dependent tasks
- Be cancelled via context

### Requirements
- Create a `Task` struct with:
  - Name (string)
  - Dependencies ([]string - names of tasks it depends on)
  - Execute func() (error)
  - Status (pending, running, completed, failed, cancelled)
- Implement a `TaskRunner` that:
  - Accepts a list of tasks with dependencies
  - Executes tasks respecting dependencies
  - Runs independent tasks concurrently
  - Propagates errors (dependent tasks skip if dependency fails)
  - Supports cancellation via context
  - Reports task status and results
- Handle circular dependencies (detect and error)
- Create a sample workflow with 6-8 tasks

### Example Workflow
```
Task A (no deps) → Task B (depends on A) → Task D (depends on B)
Task C (no deps) → Task E (depends on C)
Task F (depends on D, E)
Task G (depends on F)
```

### Expected Output
```
Starting task orchestration...

Task dependency graph:
  A (no dependencies)
  B (depends on: A)
  C (no dependencies)
  D (depends on: B)
  E (depends on: C)
  F (depends on: D, E)
  G (depends on: F)

Executing tasks...
  [STARTING] Task A
  [STARTING] Task C (concurrent with A)
  [COMPLETED] Task A (120ms)
  [STARTING] Task B
  [COMPLETED] Task C (80ms)
  [STARTING] Task E
  [COMPLETED] Task E (150ms)
  [COMPLETED] Task B (100ms)
  [STARTING] Task D
  [COMPLETED] Task D (90ms)
  [STARTING] Task F
  [FAILED] Task F: simulated error
  [SKIPPED] Task G (dependency F failed)

Orchestration Summary:
  Completed: 5/7
  Failed: 1/7
  Skipped: 1/7
  Total time: 350ms
```

### Challenge
- Add retry logic for failed tasks
- Implement task timeouts
