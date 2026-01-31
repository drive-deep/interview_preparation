# Go Concurrency

Master Go's powerful concurrency model.

## Learning Path

### Phase 1: Fundamentals (1 week)
- Goroutines
- Channels (buffered/unbuffered)
- Select statement
- sync package (Mutex, RWMutex, WaitGroup, Once)

### Phase 2: Patterns (2-3 weeks)
- Worker Pool
- Fan-In / Fan-Out
- Pipeline
- Rate Limiter
- Circuit Breaker
- Semaphore

## Folder Structure
```
go/
├── fundamentals/    # Core concepts
└── patterns/        # Concurrency patterns
```

## Key Concepts
- "Don't communicate by sharing memory; share memory by communicating"
- Always think about goroutine lifecycle
- Use `context.Context` for cancellation
- Avoid goroutine leaks
