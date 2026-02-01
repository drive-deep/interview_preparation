package abstract_factory

import "fmt"

// ============================================================================
// PRODUCT INTERFACES - Multiple related products (not just one!)
// ============================================================================

// DBConnection handles database connectivity
type DBConnection interface {
	Connect() string
	Close() string
}

// DBQuery handles SQL execution
type DBQuery interface {
	Execute(sql string) string
}

// DBTransaction handles transaction management
type DBTransaction interface {
	Begin() string
	Commit() string
	Rollback() string
}

// ============================================================================
// ABSTRACT FACTORY - Creates a FAMILY of related products
// ============================================================================

// DBFactory creates all database-related products for a specific vendor
type DBFactory interface {
	CreateConnection() DBConnection
	CreateQuery() DBQuery
	CreateTransaction() DBTransaction
}

// ============================================================================
// MYSQL FAMILY - All products designed to work together
// ============================================================================

// MySQLConnection implements DBConnection for MySQL
type MySQLConnection struct {
	url  string
	port int
}

func (m *MySQLConnection) Connect() string {
	return fmt.Sprintf("MySQL connected to %s:%d", m.url, m.port)
}

func (m *MySQLConnection) Close() string {
	return "MySQL connection closed"
}

// MySQLQuery implements DBQuery for MySQL
type MySQLQuery struct{}

func (m *MySQLQuery) Execute(sql string) string {
	return fmt.Sprintf("MySQL executing: %s", sql)
}

// MySQLTransaction implements DBTransaction for MySQL
type MySQLTransaction struct{}

func (m *MySQLTransaction) Begin() string    { return "MySQL: BEGIN TRANSACTION" }
func (m *MySQLTransaction) Commit() string   { return "MySQL: COMMIT" }
func (m *MySQLTransaction) Rollback() string { return "MySQL: ROLLBACK" }

// MySQLFactory implements DBFactory for MySQL
type MySQLFactory struct {
	url  string
	port int
}

// NewMySQLFactory creates a new MySQL factory
func NewMySQLFactory(url string, port int) *MySQLFactory {
	return &MySQLFactory{url: url, port: port}
}

func (f *MySQLFactory) CreateConnection() DBConnection {
	return &MySQLConnection{url: f.url, port: f.port}
}

func (f *MySQLFactory) CreateQuery() DBQuery {
	return &MySQLQuery{}
}

func (f *MySQLFactory) CreateTransaction() DBTransaction {
	return &MySQLTransaction{}
}

// ============================================================================
// POSTGRES FAMILY - All products designed to work together
// ============================================================================

// PostgresConnection implements DBConnection for PostgreSQL
type PostgresConnection struct {
	url  string
	port int
}

func (p *PostgresConnection) Connect() string {
	return fmt.Sprintf("Postgres connected to %s:%d", p.url, p.port)
}

func (p *PostgresConnection) Close() string {
	return "Postgres connection closed"
}

// PostgresQuery implements DBQuery for PostgreSQL
type PostgresQuery struct{}

func (p *PostgresQuery) Execute(sql string) string {
	return fmt.Sprintf("Postgres executing: %s", sql)
}

// PostgresTransaction implements DBTransaction for PostgreSQL
type PostgresTransaction struct{}

func (p *PostgresTransaction) Begin() string    { return "Postgres: BEGIN TRANSACTION" }
func (p *PostgresTransaction) Commit() string   { return "Postgres: COMMIT" }
func (p *PostgresTransaction) Rollback() string { return "Postgres: ROLLBACK" }

// PostgresFactory implements DBFactory for PostgreSQL
type PostgresFactory struct {
	url  string
	port int
}

// NewPostgresFactory creates a new Postgres factory
func NewPostgresFactory(url string, port int) *PostgresFactory {
	return &PostgresFactory{url: url, port: port}
}

func (f *PostgresFactory) CreateConnection() DBConnection {
	return &PostgresConnection{url: f.url, port: f.port}
}

func (f *PostgresFactory) CreateQuery() DBQuery {
	return &PostgresQuery{}
}

func (f *PostgresFactory) CreateTransaction() DBTransaction {
	return &PostgresTransaction{}
}

// ============================================================================
// FACTORY SELECTOR - Returns appropriate factory based on config
// ============================================================================

// GetDBFactory returns the appropriate factory for the given database type
func GetDBFactory(dbType, url string, port int) DBFactory {
	switch dbType {
	case "mysql":
		return NewMySQLFactory(url, port)
	case "postgres":
		return NewPostgresFactory(url, port)
	default:
		return NewMySQLFactory(url, port) // default to MySQL
	}
}

// ============================================================================
// CLIENT CODE - Works with any database family via interfaces only
// ============================================================================

// DatabaseClient demonstrates how client code uses the abstract factory
type DatabaseClient struct {
	factory DBFactory
	conn    DBConnection
	query   DBQuery
	tx      DBTransaction
}

// NewDatabaseClient creates a client with the given factory
func NewDatabaseClient(factory DBFactory) *DatabaseClient {
	return &DatabaseClient{
		factory: factory,
		conn:    factory.CreateConnection(),
		query:   factory.CreateQuery(),
		tx:      factory.CreateTransaction(),
	}
}

// ExecuteWithTransaction runs a query within a transaction
func (c *DatabaseClient) ExecuteWithTransaction(sql string) []string {
	results := []string{}
	results = append(results, c.conn.Connect())
	results = append(results, c.tx.Begin())
	results = append(results, c.query.Execute(sql))
	results = append(results, c.tx.Commit())
	results = append(results, c.conn.Close())
	return results
}
