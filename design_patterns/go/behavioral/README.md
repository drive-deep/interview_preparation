# Behavioral Patterns in Go

## Patterns

### 1. Strategy
- **Problem**: Swap algorithms at runtime
- **Go Approach**: Interface + DI, or function types
- **Examples**: Payment processors, compression

### 2. Observer
- **Problem**: Notify subscribers of changes
- **Go Approach**: Channels for pub/sub
- **Examples**: Event systems

### 3. Command
- **Problem**: Encapsulate requests as objects
- **Go Approach**: Functions as first-class citizens
- **Examples**: Undo/redo, task queues

### 4. State
- **Problem**: Object behavior changes with state
- **Go Approach**: Interface per state
- **Examples**: Order status, connection lifecycle

### 5. Chain of Responsibility
- **Problem**: Pass request along handler chain
- **Go Approach**: Middleware chains
- **Examples**: HTTP middleware

### 6. Template Method
- **Problem**: Define algorithm skeleton
- **Go Approach**: Higher-order functions
- **Examples**: HTTP handlers

### 7. Iterator
- **Problem**: Sequential access without exposing internals
- **Go Approach**: Channels, `Next()` method
- **Examples**: Collection traversal

## Practice Order
Strategy → Observer → Command → State → Chain of Responsibility
