package builder

import (
	"testing"
	"time"
)

// TODO: Write tests for Builder pattern
//
// ==================== POSITIVE TEST CASES ====================
//
// func TestServerBuilder_DefaultValues(t *testing.T) {
//     // Build with no options, verify defaults
// }
//
// func TestServerBuilder_AllOptions(t *testing.T) {
//     // Build with all options set
// }
//
// func TestServerBuilder_PartialOptions(t *testing.T) {
//     // Build with only some options
// }
//
// func TestServerBuilder_MethodChaining(t *testing.T) {
//     // Verify method chaining returns same builder
// }
//
// ==================== NEGATIVE TEST CASES ====================
//
// func TestServerBuilder_InvalidPort(t *testing.T) {
//     // Port < 0 or > 65535 should error
// }
//
// func TestServerBuilder_EmptyHost(t *testing.T) {
//     // Empty host should error or use default
// }
//
// func TestServerBuilder_NegativeTimeout(t *testing.T) {
//     // Negative timeout should error
// }
//
// ==================== FUNCTIONAL OPTIONS TESTS ====================
//
// func TestNewServer_WithOptions(t *testing.T) {
//     server := NewServer(WithHost("test"), WithPort(9000))
//     // verify options applied
// }
//
// func TestNewServer_NoOptions(t *testing.T) {
//     server := NewServer()
//     // verify defaults
// }


func TestServerBuilder_DefaultValues(t *testing.T) {
	builder := NewServerBuilder()
	server := builder.Build()

	if server.host != "localhost" {
		t.Errorf("Expected host %q, got %q", "localhost", server.host)
	}
	if server.port != 8080 {
		t.Errorf("Expected port %d, got %d", 8080, server.port)
	}
	if server.tls != false {
		t.Errorf("Expected TLS %t, got %t", false, server.tls)
	}
	if server.timeout != 30*time.Second {
		t.Errorf("Expected timeout %s, got %s", 30*time.Second, server.timeout)
	}
	if server.maxConnections != 100 {
		t.Errorf("Expected max connections %d, got %d", 100, server.maxConnections)
	}
}

func TestServerBuilder_AllOptions(t *testing.T) {
	builder := NewServerBuilder()
	server := builder.WithHost("api.example.com").WithPort(443).WithTLS(true).WithTimeout(10*time.Second).WithMaxConnections(500).Build()

	if server.host != "api.example.com" {
		t.Errorf("Expected host %q, got %q", "api.example.com", server.host)
	}
	if server.port != 443 {
		t.Errorf("Expected port %d, got %d", 443, server.port)
	}
	if server.tls != true {
		t.Errorf("Expected TLS %t, got %t", true, server.tls)
	}
	if server.timeout != 10*time.Second {
		t.Errorf("Expected timeout %s, got %s", 10*time.Second, server.timeout)
	}
	if server.maxConnections != 500 {
		t.Errorf("Expected max connections %d, got %d", 500, server.maxConnections)
	}
}

func TestServerBuilder_PartialOptions(t *testing.T) {
	builder := NewServerBuilder()
	server := builder.WithHost("api.example.com").WithPort(443).Build()

	if server.host != "api.example.com" {
		t.Errorf("Expected host %q, got %q", "api.example.com", server.host)
	}
	if server.port != 443 {
		t.Errorf("Expected port %d, got %d", 443, server.port)
	}
	if server.tls != false {
		t.Errorf("Expected TLS %t, got %t", false, server.tls)
	}
	if server.timeout != 30*time.Second {
		t.Errorf("Expected timeout %s, got %s", 30*time.Second, server.timeout)
	}
	if server.maxConnections != 100 {
		t.Errorf("Expected max connections %d, got %d", 100, server.maxConnections)
	}
}


func TestServerBuilder_MethodChaining(t *testing.T) {
	builder := NewServerBuilder()
	server := builder.WithHost("api.example.com").WithPort(443).WithTLS(true).WithTimeout(10*time.Second).WithMaxConnections(500).Build()

	if server.host != "api.example.com" {
		t.Errorf("Expected host %q, got %q", "api.example.com", server.host)
	}
	if server.port != 443 {
		t.Errorf("Expected port %d, got %d", 443, server.port)
	}
	if server.tls != true {
		t.Errorf("Expected TLS %t, got %t", true, server.tls)
	}
	if server.timeout != 10*time.Second {
		t.Errorf("Expected timeout %s, got %s", 10*time.Second, server.timeout)
	}
	if server.maxConnections != 500 {
		t.Errorf("Expected max connections %d, got %d", 500, server.maxConnections)
	}
}

	