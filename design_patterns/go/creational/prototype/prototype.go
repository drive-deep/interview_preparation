package prototype

// ============================================================================
// PROTOTYPE INTERFACE
// ============================================================================

// Prototype defines the cloning interface
type Prototype interface {
	Clone() Prototype
}

// ============================================================================
// EXAMPLE 1: Document Prototype
// ============================================================================

// Document represents a document that can be cloned
type Document struct {
	Title    string
	Content  string
	Author   string
	Metadata map[string]string
}

// Clone creates a deep copy of the Document
func (d *Document) Clone() Prototype {
	// Deep copy the metadata map
	newMetadata := make(map[string]string)
	for k, v := range d.Metadata {
		newMetadata[k] = v
	}

	return &Document{
		Title:    d.Title,
		Content:  d.Content,
		Author:   d.Author,
		Metadata: newMetadata,
	}
}

// ============================================================================
// EXAMPLE 2: Server Config Prototype
// ============================================================================

// ServerConfig represents a server configuration that can be cloned
type ServerConfig struct {
	Host       string
	Port       int
	TLS        bool
	Timeout    int
	AllowedIPs []string
	Headers    map[string]string
}

// Clone creates a deep copy of the ServerConfig
func (s *ServerConfig) Clone() Prototype {
	// Deep copy the AllowedIPs slice
	newAllowedIPs := make([]string, len(s.AllowedIPs))
	copy(newAllowedIPs, s.AllowedIPs)

	// Deep copy the Headers map
	newHeaders := make(map[string]string)
	for k, v := range s.Headers {
		newHeaders[k] = v
	}

	return &ServerConfig{
		Host:       s.Host,
		Port:       s.Port,
		TLS:        s.TLS,
		Timeout:    s.Timeout,
		AllowedIPs: newAllowedIPs,
		Headers:    newHeaders,
	}
}

// ============================================================================
// PROTOTYPE REGISTRY - Centralized storage for prototypes
// ============================================================================

// Registry stores prototype instances by name
type Registry struct {
	prototypes map[string]Prototype
}

// NewRegistry creates a new prototype registry
func NewRegistry() *Registry {
	return &Registry{
		prototypes: make(map[string]Prototype),
	}
}

// Register adds a prototype to the registry
func (r *Registry) Register(name string, proto Prototype) {
	r.prototypes[name] = proto
}

// Get retrieves and clones a prototype by name
// Returns nil if prototype not found
func (r *Registry) Get(name string) Prototype {
	proto, exists := r.prototypes[name]
	if !exists {
		return nil
	}
	// Always return a CLONE, never the original!
	return proto.Clone()
}

// Has checks if a prototype exists in the registry
func (r *Registry) Has(name string) bool {
	_, exists := r.prototypes[name]
	return exists
}

// Unregister removes a prototype from the registry
func (r *Registry) Unregister(name string) {
	delete(r.prototypes, name)
}
