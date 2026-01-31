# Go Concurrency Fundamentals

## Topics to Master

### 1. Goroutines
- Lightweight threads managed by Go runtime
- Stack starts at 2KB (grows as needed)
- Scheduled by Go scheduler (M:N threading)

### 2. Channels
- **Unbuffered**: Synchronous, blocks until receiver ready
- **Buffered**: Async up to capacity
- **Directional**: `chan<-` (send-only), `<-chan` (receive-only)
- Close channels when done sending

### 3. Select Statement
- Multiplex channel operations
- `default` case for non-blocking
- Random selection when multiple ready

### 4. sync Package
| Primitive | Use Case |
|-----------|----------|
| `Mutex` | Protect shared state |
| `RWMutex` | Multiple readers, single writer |
| `WaitGroup` | Wait for goroutines to finish |
| `Once` | Run initialization once |
| `Cond` | Wait for conditions |
| `Pool` | Reuse objects |

### 5. Context
- Cancellation propagation
- Timeouts and deadlines
- Request-scoped values

## Practice Exercises
1. Producer-consumer with channels
2. Timeout with select
3. Safe counter with mutex
4. Graceful shutdown with context
