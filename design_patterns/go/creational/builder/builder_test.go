package builder

import "testing"

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
