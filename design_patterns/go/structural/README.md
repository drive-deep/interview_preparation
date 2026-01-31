# Structural Patterns in Go

## Patterns

### 1. Adapter
- **Problem**: Make incompatible interfaces work together
- **Go Approach**: Wrapper struct implementing target interface
- **Examples**: Legacy API integration

### 2. Decorator
- **Problem**: Add behavior without modifying original
- **Go Approach**: Middleware pattern
- **Examples**: HTTP middleware (logging, auth)

### 3. Facade
- **Problem**: Simplified interface to complex subsystem
- **Go Approach**: Package-level functions
- **Examples**: High-level DB API

### 4. Composite
- **Problem**: Treat individual and composite objects uniformly
- **Go Approach**: Recursive interface
- **Examples**: File system, UI components

### 5. Proxy
- **Problem**: Control access to another object
- **Go Approach**: Wrapper with same interface
- **Examples**: Caching proxy, lazy loading

### 6. Bridge
- **Problem**: Separate abstraction from implementation
- **Go Approach**: Composition with interfaces
- **Examples**: Multiple DB backends

## Practice Order
Adapter → Decorator → Facade → Composite → Proxy → Bridge
