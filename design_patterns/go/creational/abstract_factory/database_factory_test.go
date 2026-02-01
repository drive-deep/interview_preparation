package abstract_factory

import (
	"strings"
	"testing"
)

// ============================================================================
// MYSQL FACTORY TESTS
// ============================================================================

func TestMySQLFactory_CreateConnection(t *testing.T) {
	factory := NewMySQLFactory("localhost", 3306)
	conn := factory.CreateConnection()

	result := conn.Connect()
	if !strings.Contains(result, "MySQL") {
		t.Errorf("Expected MySQL connection, got: %s", result)
	}
	if !strings.Contains(result, "localhost:3306") {
		t.Errorf("Expected localhost:3306, got: %s", result)
	}
}

func TestMySQLFactory_CreateQuery(t *testing.T) {
	factory := NewMySQLFactory("localhost", 3306)
	query := factory.CreateQuery()

	result := query.Execute("SELECT * FROM users")
	if !strings.Contains(result, "MySQL") {
		t.Errorf("Expected MySQL query, got: %s", result)
	}
	if !strings.Contains(result, "SELECT * FROM users") {
		t.Errorf("Expected SQL in result, got: %s", result)
	}
}

func TestMySQLFactory_CreateTransaction(t *testing.T) {
	factory := NewMySQLFactory("localhost", 3306)
	tx := factory.CreateTransaction()

	if !strings.Contains(tx.Begin(), "MySQL") {
		t.Error("Expected MySQL transaction begin")
	}
	if !strings.Contains(tx.Commit(), "MySQL") {
		t.Error("Expected MySQL transaction commit")
	}
	if !strings.Contains(tx.Rollback(), "MySQL") {
		t.Error("Expected MySQL transaction rollback")
	}
}

// ============================================================================
// POSTGRES FACTORY TESTS
// ============================================================================

func TestPostgresFactory_CreateConnection(t *testing.T) {
	factory := NewPostgresFactory("db.example.com", 5432)
	conn := factory.CreateConnection()

	result := conn.Connect()
	if !strings.Contains(result, "Postgres") {
		t.Errorf("Expected Postgres connection, got: %s", result)
	}
	if !strings.Contains(result, "db.example.com:5432") {
		t.Errorf("Expected db.example.com:5432, got: %s", result)
	}
}

func TestPostgresFactory_CreateQuery(t *testing.T) {
	factory := NewPostgresFactory("localhost", 5432)
	query := factory.CreateQuery()

	result := query.Execute("INSERT INTO orders VALUES (1)")
	if !strings.Contains(result, "Postgres") {
		t.Errorf("Expected Postgres query, got: %s", result)
	}
}

func TestPostgresFactory_CreateTransaction(t *testing.T) {
	factory := NewPostgresFactory("localhost", 5432)
	tx := factory.CreateTransaction()

	if !strings.Contains(tx.Begin(), "Postgres") {
		t.Error("Expected Postgres transaction begin")
	}
	if !strings.Contains(tx.Commit(), "Postgres") {
		t.Error("Expected Postgres transaction commit")
	}
	if !strings.Contains(tx.Rollback(), "Postgres") {
		t.Error("Expected Postgres transaction rollback")
	}
}

// ============================================================================
// FACTORY SELECTOR TESTS
// ============================================================================

func TestGetDBFactory_MySQL(t *testing.T) {
	factory := GetDBFactory("mysql", "localhost", 3306)
	conn := factory.CreateConnection()

	if !strings.Contains(conn.Connect(), "MySQL") {
		t.Error("GetDBFactory(mysql) should return MySQL factory")
	}
}

func TestGetDBFactory_Postgres(t *testing.T) {
	factory := GetDBFactory("postgres", "localhost", 5432)
	conn := factory.CreateConnection()

	if !strings.Contains(conn.Connect(), "Postgres") {
		t.Error("GetDBFactory(postgres) should return Postgres factory")
	}
}

func TestGetDBFactory_Default(t *testing.T) {
	factory := GetDBFactory("unknown", "localhost", 1234)
	conn := factory.CreateConnection()

	// Should default to MySQL
	if !strings.Contains(conn.Connect(), "MySQL") {
		t.Error("GetDBFactory(unknown) should default to MySQL factory")
	}
}

// ============================================================================
// CLIENT INTEGRATION TESTS
// ============================================================================

func TestDatabaseClient_MySQL(t *testing.T) {
	factory := NewMySQLFactory("localhost", 3306)
	client := NewDatabaseClient(factory)

	results := client.ExecuteWithTransaction("SELECT * FROM users")

	// Verify all operations are MySQL
	for _, result := range results {
		if !strings.Contains(result, "MySQL") {
			t.Errorf("Expected all operations to be MySQL, got: %s", result)
		}
	}

	// Verify correct order: Connect -> Begin -> Execute -> Commit -> Close
	if len(results) != 5 {
		t.Errorf("Expected 5 operations, got %d", len(results))
	}
}

func TestDatabaseClient_Postgres(t *testing.T) {
	factory := NewPostgresFactory("localhost", 5432)
	client := NewDatabaseClient(factory)

	results := client.ExecuteWithTransaction("INSERT INTO logs VALUES (1)")

	// Verify all operations are Postgres
	for _, result := range results {
		if !strings.Contains(result, "Postgres") {
			t.Errorf("Expected all operations to be Postgres, got: %s", result)
		}
	}
}

// ============================================================================
// FAMILY CONSISTENCY TEST - Key benefit of Abstract Factory!
// ============================================================================

func TestFactoryProducesConsistentDBFamily(t *testing.T) {
	// MySQL family - all products should be MySQL
	mysqlFactory := NewMySQLFactory("localhost", 3306)
	mysqlConn := mysqlFactory.CreateConnection()
	mysqlQuery := mysqlFactory.CreateQuery()
	mysqlTx := mysqlFactory.CreateTransaction()

	if !strings.Contains(mysqlConn.Connect(), "MySQL") ||
		!strings.Contains(mysqlQuery.Execute("test"), "MySQL") ||
		!strings.Contains(mysqlTx.Begin(), "MySQL") {
		t.Error("MySQL factory should produce consistent MySQL-family products")
	}

	// Postgres family - all products should be Postgres
	pgFactory := NewPostgresFactory("localhost", 5432)
	pgConn := pgFactory.CreateConnection()
	pgQuery := pgFactory.CreateQuery()
	pgTx := pgFactory.CreateTransaction()

	if !strings.Contains(pgConn.Connect(), "Postgres") ||
		!strings.Contains(pgQuery.Execute("test"), "Postgres") ||
		!strings.Contains(pgTx.Begin(), "Postgres") {
		t.Error("Postgres factory should produce consistent Postgres-family products")
	}
}

// ============================================================================
// SWITCHING FACTORIES TEST - Demonstrates the power of Abstract Factory
// ============================================================================

func TestSwitchingFactories(t *testing.T) {
	runOperations := func(factory DBFactory) string {
		conn := factory.CreateConnection()
		return conn.Connect()
	}

	// Start with MySQL
	result1 := runOperations(NewMySQLFactory("localhost", 3306))
	if !strings.Contains(result1, "MySQL") {
		t.Error("Should be using MySQL")
	}

	// Switch to Postgres - same client code, different factory
	result2 := runOperations(NewPostgresFactory("localhost", 5432))
	if !strings.Contains(result2, "Postgres") {
		t.Error("Should be using Postgres after switch")
	}
}
