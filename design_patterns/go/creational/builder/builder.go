package builder


import (
	"time"
	"fmt"
)


type Server struct {
	host             string
	port             int
	tls              bool
	timeout          time.Duration
	maxConnections int
}

func NewServerBuilder() *Server {
	return &Server{
		host: "localhost",
		port: 8080,
		tls: false,
		timeout: 30 * time.Second,
		maxConnections: 100,
	}
}

func (b *Server) WithHost(host string) *Server {
	b.host = host
	return b
}

func (b *Server) WithPort(port int) *Server {
	b.port = port
	return b
}

func (b *Server) WithTLS(tls bool) *Server {
	b.tls = tls
	return b
}

func (b *Server) WithTimeout(timeout time.Duration) *Server {
	b.timeout = timeout
	return b
}

func (b *Server) WithMaxConnections(maxConnections int) *Server {
	b.maxConnections = maxConnections
	return b
}

func (b *Server) Build() *Server {
	return b
}

func (s *Server) Start() {
	fmt.Printf("Starting server on %s:%d with TLS: %t, timeout: %s, max connections: %d\n", s.host, s.port, s.tls, s.timeout, s.maxConnections)
}	

func main() {
	server := NewServerBuilder().WithHost("api.example.com").WithPort(443).WithTLS(true).WithTimeout(10*time.Second).WithMaxConnections(500).Build()
	server.Start()
}

