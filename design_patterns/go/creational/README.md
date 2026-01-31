# Creational Patterns in Go

## Patterns

### 1. Singleton
- **Problem**: Ensure only one instance exists
- **Go Approach**: `sync.Once` for thread-safe lazy init
- **Examples**: Logger, Config, DB pool

### 2. Factory Method
- **Problem**: Create objects without specifying exact type
- **Go Approach**: Functions returning interfaces
- **Examples**: Database drivers, parsers

### 3. Builder
- **Problem**: Construct complex objects step-by-step
- **Go Approach**: Functional options pattern
- **Examples**: HTTP client config, query builders

### 4. Prototype
- **Problem**: Clone existing objects
- **Go Approach**: `Clone()` method with deep copy
- **Examples**: Template objects, configs

### 5. Abstract Factory
- **Problem**: Create families of related objects
- **Go Approach**: Factory interface
- **Examples**: UI toolkit for different OS

## Practice Order
Singleton → Factory → Builder → Prototype → Abstract Factory
