# Builder Pattern

## Intent
Separate the construction of a complex object from its representation, allowing the same construction process to create different representations.

---

## 🎤 Interview Questions & Answers

### Q1: What is the Builder Pattern?
**A:** Builder Pattern is a creational pattern that constructs complex objects step-by-step. Instead of a constructor with many parameters, we use a builder that sets each property individually and returns the final object.

### Q2: Why do we use Builder Pattern?
**A:** We use Builder when:
- Object has **many optional parameters** (constructor would be messy)
- We want to create **immutable objects** step-by-step
- We need **different representations** of the same object
- We want **readable, self-documenting** code

**Example:** Building an HTTP server config with optional TLS, timeout, max connections, etc.

### Q3: What problem does Builder Pattern solve?
**A:** It solves the **telescoping constructor** problem.

**Without Builder (Bad):**
```go
// Which parameter is which? Hard to read!
server := NewServer("localhost", 8080, true, 30, 100, nil, "v1")
```

**With Builder (Good):**
```go
server := NewServerBuilder().
    WithHost("localhost").
    WithPort(8080).
    WithTLS(true).
    WithTimeout(30).
    Build()
```

### Q4: What is Go's Functional Options Pattern?
**A:** It's Go's idiomatic way to implement Builder. Instead of method chaining, we use variadic functions:

```go
server := NewServer(
    WithHost("localhost"),
    WithPort(8080),
    WithTLS(),
)
```

Each `With...` function returns an `Option` type that modifies the config.

### Q5: Builder vs Factory - What's the difference?
**A:**
| Builder | Factory |
|---------|---------|
| Creates ONE complex object step-by-step | Creates different types of objects |
| Focus: construction process | Focus: which type to create |
| Returns same type with different configs | Returns different concrete types |
| `ServerBuilder.WithPort(8080).Build()` | `NotifierFactory("email")` |

### Q6: When should I NOT use Builder Pattern?
**A:** Don't use it when:
- Object has few parameters (just use constructor)
- All parameters are required (no optional ones)
- Object creation is simple
- Adds unnecessary complexity

### Q7: Give real-world examples of Builder Pattern in Go.
**A:**
- **`http.Client`**: Build with custom transport, timeout, redirect policy
- **`grpc.Dial`**: `grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())`
- **`strings.Builder`**: Build strings efficiently
- **SQL query builders**: `query.Select("*").From("users").Where("id = ?", 1)`

---

## Structure (Traditional Builder)

```
┌─────────────────────────────────┐
│          ServerBuilder          │
├─────────────────────────────────┤
│ - host: string                  │
│ - port: int                     │
│ - tls: bool                     │
├─────────────────────────────────┤
│ + WithHost(h) *ServerBuilder    │
│ + WithPort(p) *ServerBuilder    │
│ + WithTLS() *ServerBuilder      │
│ + Build() *Server               │
└─────────────────────────────────┘
          │
          ▼ creates
┌─────────────────────────────────┐
│            Server               │
└─────────────────────────────────┘
```

## Structure (Go Functional Options)

```
┌─────────────────────────────────┐
│   type Option func(*Config)     │
└─────────────────────────────────┘

WithHost("x")  ──┐
WithPort(8080) ──┼──> NewServer(opts...) ──> *Server
WithTLS()      ──┘
```

---

## Two Approaches in Go

### Approach 1: Traditional Builder (Method Chaining)
```go
server := NewServerBuilder().
    WithHost("localhost").
    WithPort(8080).
    Build()
```

### Approach 2: Functional Options (Go Idiom) ⭐
```go
server := NewServer(
    WithHost("localhost"),
    WithPort(8080),
)
```

**Recommended:** Functional Options is more Go-idiomatic.

---

## Practice Exercise

Implement an **HTTP Server Builder** with:
- `host` (default: "localhost")
- `port` (default: 8080)
- `tls` (default: false)
- `timeout` (default: 30s)
- `maxConnections` (default: 100)

Try BOTH approaches:
1. Traditional Builder with method chaining
2. Functional Options pattern

---

## Test Cases to Write

### Positive Tests
1. Build with default values
2. Build with all options set
3. Build with partial options
4. Method chaining returns same builder

### Negative Tests
5. Invalid port (negative, > 65535)
6. Empty host validation
7. Negative timeout

---

## Files
- `builder.go` - Your implementation
- `builder_test.go` - Your tests
