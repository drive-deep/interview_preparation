# Design Patterns in Go

Go-idiomatic implementations of GoF patterns.

## Key Go Idioms

| Pattern | Java/OOP Way | Go Idiomatic Way |
|---------|--------------|------------------|
| Singleton | Class with private constructor | `sync.Once` + package-level var |
| Factory | Abstract factory class | Functions returning interfaces |
| Builder | Builder class with methods | Functional options (`With...`) |
| Decorator | Inheritance chain | Middleware, wrapper functions |
| Observer | Observer interface | Channels |
| Strategy | Strategy interface | Interface + DI, or function types |

## Folder Structure

```
go/
├── creational/     # Singleton, Factory, Builder, Prototype
├── structural/     # Adapter, Decorator, Facade, Composite
└── behavioral/     # Strategy, Observer, Command, State
```

## Practice Tips

- Use interfaces extensively (implicit in Go)
- Prefer composition over inheritance
- Think about concurrency safety from the start
