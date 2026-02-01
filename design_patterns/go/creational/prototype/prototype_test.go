package prototype

import (
	"testing"
)

// ============================================================================
// DOCUMENT CLONE TESTS
// ============================================================================

func TestDocument_Clone_Basic(t *testing.T) {
	original := &Document{
		Title:   "Report Template",
		Content: "This is the base content",
		Author:  "System",
	}

	clone := original.Clone().(*Document)

	// Verify values are copied
	if clone.Title != original.Title {
		t.Errorf("Expected title %s, got %s", original.Title, clone.Title)
	}
	if clone.Content != original.Content {
		t.Errorf("Expected content %s, got %s", original.Content, clone.Content)
	}
	if clone.Author != original.Author {
		t.Errorf("Expected author %s, got %s", original.Author, clone.Author)
	}

	// Modify clone - should NOT affect original
	clone.Author = "Prabhat"
	if original.Author == "Prabhat" {
		t.Error("Modifying clone should NOT affect original")
	}
}

func TestDocument_Clone_DeepCopy_Metadata(t *testing.T) {
	original := &Document{
		Title:   "Config Doc",
		Content: "...",
		Author:  "Admin",
		Metadata: map[string]string{
			"version": "1.0",
			"type":    "internal",
		},
	}

	clone := original.Clone().(*Document)

	// Verify metadata is copied
	if clone.Metadata["version"] != "1.0" {
		t.Error("Metadata should be copied")
	}

	// Modify clone's metadata - should NOT affect original
	clone.Metadata["version"] = "2.0"
	if original.Metadata["version"] == "2.0" {
		t.Error("Deep copy failed: modifying clone's map affected original")
	}
}

func TestDocument_Clone_NilMetadata(t *testing.T) {
	original := &Document{
		Title:   "Simple Doc",
		Content: "No metadata",
		Author:  "User",
	}

	clone := original.Clone().(*Document)

	if clone.Title != original.Title {
		t.Error("Clone should copy basic fields even with nil metadata")
	}
}

// ============================================================================
// SERVER CONFIG CLONE TESTS
// ============================================================================

func TestServerConfig_Clone_Basic(t *testing.T) {
	original := &ServerConfig{
		Host:    "localhost",
		Port:    8080,
		TLS:     true,
		Timeout: 30,
	}

	clone := original.Clone().(*ServerConfig)

	if clone.Host != original.Host || clone.Port != original.Port {
		t.Error("Clone should have same values as original")
	}
	if clone.TLS != original.TLS || clone.Timeout != original.Timeout {
		t.Error("Clone should have same values as original")
	}

	// Modify clone
	clone.Port = 9090
	if original.Port == 9090 {
		t.Error("Modifying clone should NOT affect original")
	}
}

func TestServerConfig_Clone_DeepCopy_Slice(t *testing.T) {
	original := &ServerConfig{
		Host:       "api.example.com",
		Port:       443,
		AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
	}

	clone := original.Clone().(*ServerConfig)

	// Verify slice is copied
	if len(clone.AllowedIPs) != 2 {
		t.Error("AllowedIPs should be copied")
	}

	// Modify clone's slice - should NOT affect original
	clone.AllowedIPs[0] = "0.0.0.0"
	if original.AllowedIPs[0] == "0.0.0.0" {
		t.Error("Deep copy failed: modifying clone's slice affected original")
	}
}

func TestServerConfig_Clone_DeepCopy_Map(t *testing.T) {
	original := &ServerConfig{
		Host: "localhost",
		Port: 8080,
		Headers: map[string]string{
			"X-Api-Key":    "secret",
			"Content-Type": "application/json",
		},
	}

	clone := original.Clone().(*ServerConfig)

	// Verify map is copied
	if clone.Headers["X-Api-Key"] != "secret" {
		t.Error("Headers should be copied")
	}

	// Modify clone's map - should NOT affect original
	clone.Headers["X-Api-Key"] = "new-secret"
	if original.Headers["X-Api-Key"] == "new-secret" {
		t.Error("Deep copy failed: modifying clone's map affected original")
	}
}

// ============================================================================
// REGISTRY TESTS
// ============================================================================

func TestRegistry_RegisterAndGet(t *testing.T) {
	registry := NewRegistry()

	baseDoc := &Document{
		Title:   "Base Template",
		Content: "Default content",
		Author:  "System",
	}

	registry.Register("default-doc", baseDoc)

	// Get should return a CLONE, not the original
	clone := registry.Get("default-doc").(*Document)
	clone.Author = "Modified"

	// Original in registry should be unchanged
	original := registry.Get("default-doc").(*Document)
	if original.Author == "Modified" {
		t.Error("Registry.Get should return clones, not originals")
	}
}

func TestRegistry_GetNonExistent(t *testing.T) {
	registry := NewRegistry()

	result := registry.Get("does-not-exist")
	if result != nil {
		t.Error("Get on non-existent key should return nil")
	}
}

func TestRegistry_Has(t *testing.T) {
	registry := NewRegistry()

	registry.Register("my-config", &ServerConfig{Host: "localhost"})

	if !registry.Has("my-config") {
		t.Error("Has should return true for registered prototype")
	}
	if registry.Has("unknown") {
		t.Error("Has should return false for unregistered prototype")
	}
}

func TestRegistry_Unregister(t *testing.T) {
	registry := NewRegistry()

	registry.Register("temp-doc", &Document{Title: "Temp"})
	registry.Unregister("temp-doc")

	if registry.Has("temp-doc") {
		t.Error("Unregister should remove the prototype")
	}
}

func TestRegistry_MultiplePrototypes(t *testing.T) {
	registry := NewRegistry()

	registry.Register("invoice", &Document{Title: "Invoice", Author: "Finance"})
	registry.Register("report", &Document{Title: "Report", Author: "Management"})
	registry.Register("prod-server", &ServerConfig{Host: "prod.example.com", Port: 443})

	invoice := registry.Get("invoice").(*Document)
	report := registry.Get("report").(*Document)
	server := registry.Get("prod-server").(*ServerConfig)

	if invoice.Title != "Invoice" || report.Title != "Report" {
		t.Error("Registry should store and retrieve multiple prototypes")
	}
	if server.Host != "prod.example.com" {
		t.Error("Registry should handle different prototype types")
	}
}

// ============================================================================
// PRACTICAL USE CASE TEST
// ============================================================================

func TestPrototype_PracticalExample(t *testing.T) {
	// Create a production server config template
	prodTemplate := &ServerConfig{
		Host:       "api.company.com",
		Port:       443,
		TLS:        true,
		Timeout:    30,
		AllowedIPs: []string{"10.0.0.0/8"},
		Headers: map[string]string{
			"X-Environment": "production",
			"X-Rate-Limit":  "1000",
		},
	}

	// Clone it for a new microservice
	userService := prodTemplate.Clone().(*ServerConfig)
	userService.Headers["X-Service"] = "user-service"

	// Clone for another microservice
	orderService := prodTemplate.Clone().(*ServerConfig)
	orderService.Headers["X-Service"] = "order-service"
	orderService.Headers["X-Rate-Limit"] = "500" // Different rate limit

	// Verify each clone is independent
	if prodTemplate.Headers["X-Service"] != "" {
		t.Error("Original template should NOT have X-Service header")
	}
	if userService.Headers["X-Service"] != "user-service" {
		t.Error("User service should have its own header")
	}
	if orderService.Headers["X-Rate-Limit"] != "500" {
		t.Error("Order service should have modified rate limit")
	}
	if prodTemplate.Headers["X-Rate-Limit"] != "1000" {
		t.Error("Original template rate limit should be unchanged")
	}
}
