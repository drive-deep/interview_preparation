# Singleton Pattern

## Intent
Ensure a class has only one instance and provide a global point of access to it.

## Problem It Solves
- Need exactly one instance of a class (logger, config, connection pool)
- Need global access point without global variables
- Need lazy initialization with thread-safety

## Go-Idiomatic Approach

| Java/OOP Way | Go Way |
|--------------|--------|
| Private constructor | Unexported struct (lowercase) |
| Static instance | Package-level variable |
| Synchronized block | `sync.Once` |
| getInstance() | `GetInstance()` function |

## Structure

```
┌─────────────────────────────────┐
│         Singleton               │
├─────────────────────────────────┤
│ - instance: *singleton          │
│ - once: sync.Once               │
├─────────────────────────────────┤
│ + GetInstance(): *Singleton     │
│ + DoSomething()                 │
└─────────────────────────────────┘
```

## Key Components

| Component | Purpose |
|-----------|---------|
| `sync.Once` | Ensures initialization runs exactly once (thread-safe) |
| Unexported struct | Prevents `singleton{}` instantiation from outside |
| Exported getter | Single controlled access point |

## When to Use
- ✅ Logger systems
- ✅ Configuration managers
- ✅ Database connection pools
- ✅ Cache managers

## When NOT to Use
- ❌ When you need multiple instances
- ❌ When global state causes testing issues
- ❌ When dependency injection is cleaner

## Test Cases to Write
1. Same instance returned on multiple calls
2. Thread-safety with concurrent goroutines
3. Instance is properly initialized
4. State persists across calls

## Practice Exercise
Implement a **Logger** singleton with:
- `Info(msg string)`
- `Error(msg string)`
- `GetLogCount() int`

## Files
- `singleton.go` - Your implementation
- `singleton_test.go` - Your tests
