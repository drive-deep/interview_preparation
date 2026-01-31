# Go Concurrency Patterns

## Pattern Catalog

### 1. Worker Pool
- Fixed number of workers processing jobs from queue
- Controls resource usage
- **Use**: Batch processing, rate limiting

### 2. Fan-Out / Fan-In
- **Fan-Out**: Multiple goroutines read from same channel
- **Fan-In**: Multiple channels merged into one
- **Use**: Parallel processing stages

### 3. Pipeline
- Chain of processing stages
- Each stage: goroutines + channels
- **Use**: Data transformation, ETL

### 4. Rate Limiter
- Control request rate using `time.Ticker`
- Token bucket algorithm
- **Use**: API rate limiting

### 5. Circuit Breaker
- Fail fast when service is down
- States: Closed → Open → Half-Open
- **Use**: Resilient service calls

### 6. Semaphore
- Limit concurrent operations
- Buffered channel as semaphore
- **Use**: Connection pools, resource limits

### 7. Pub/Sub
- Publishers send to topic
- Subscribers receive from topic
- **Use**: Event-driven systems

### 8. Or-Done Channel
- Cancel pipeline on first error
- Merge done channels
- **Use**: Graceful shutdown

## Practice Order
Worker Pool → Fan-Out/In → Pipeline → Rate Limiter → Circuit Breaker
