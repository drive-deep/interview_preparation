package builder

// TODO: Implement Builder pattern
//
// Practice Exercise: HTTP Server Builder
//
// Fields to configure:
// - host (default: "localhost")
// - port (default: 8080)
// - tls (default: false)
// - timeout (default: 30s)
// - maxConnections (default: 100)
//
// ==================== APPROACH 1: Traditional Builder ====================
//
// type ServerBuilder struct {
//     host string
//     port int
//     // ...
// }
//
// func NewServerBuilder() *ServerBuilder {
//     return &ServerBuilder{
//         host: "localhost",
//         port: 8080,
//         // defaults...
//     }
// }
//
// func (b *ServerBuilder) WithHost(host string) *ServerBuilder {
//     b.host = host
//     return b  // return self for chaining
// }
//
// func (b *ServerBuilder) Build() *Server {
//     return &Server{host: b.host, port: b.port, ...}
// }
//
// Usage:
//   server := NewServerBuilder().WithHost("api.example.com").WithPort(443).Build()
//
// ==================== APPROACH 2: Functional Options (Go Idiom) ====================
//
// type Option func(*Server)
//
// func WithHost(host string) Option {
//     return func(s *Server) {
//         s.host = host
//     }
// }
//
// func NewServer(opts ...Option) *Server {
//     s := &Server{host: "localhost", port: 8080}  // defaults
//     for _, opt := range opts {
//         opt(s)
//     }
//     return s
// }
//
// Usage:
//   server := NewServer(WithHost("api.example.com"), WithPort(443))
