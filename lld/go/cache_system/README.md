# Cache System (LRU) - Go

## Requirements
- Get/Put operations
- LRU eviction policy
- Configurable capacity
- Thread-safe

## Patterns
- **Strategy**: Eviction policies (LRU, LFU, FIFO)
- **Decorator**: Add logging, metrics
- **Singleton**: Global cache instance

## Data Structures
- Doubly linked list + hash map
- O(1) get and put

## Concurrency
- `sync.RWMutex` for read/write locks
- Consider sharding for high concurrency

## Entities
Cache, CacheItem, EvictionPolicy
