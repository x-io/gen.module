package dbs

import (
	"database/sql"
	"time"

	//pg
	_ "github.com/lib/pq"
)

var db *sql.DB
var useSQLite bool

// Init connects to the database.
func Init(uri string) error {
	var err error
	db, err = sql.Open("postgres", uri)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func QueryRow(query string, args ...interface{}) *sql.Row {
	return db.QueryRow(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

// Prepare creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's Close method
// when the statement is no longer needed.
func Prepare(query string) (*sql.Stmt, error) {
	return db.Prepare(query)
}

// Begin starts a transaction. The default isolation level is dependent on
// the driver.
func Begin() (*sql.Tx, error) {
	return db.Begin()
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func Ping() error {
	return db.Ping()
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server
// to finish.
//
// It is rare to Close a DB, as the DB handle is meant to be
// long-lived and shared between many goroutines.
func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// SetMaxIdleConns sets the maximum number of connections in the idle
// connection pool.
//
// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
//
// If n <= 0, no idle connections are retained.
//
// The default max idle connections is currently 2. This may change in
// a future release.
func SetMaxIdleConns(n int) {
	db.SetMaxIdleConns(n)
}

// SetMaxOpenConns sets the maximum number of open connections to the database.
//
// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
// MaxIdleConns, then MaxIdleConns will be reduced to match the new
// MaxOpenConns limit.
//
// If n <= 0, then there is no limit on the number of open connections.
// The default is 0 (unlimited).
func SetMaxOpenConns(n int) {
	db.SetMaxOpenConns(n)
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are reused forever.
func SetConnMaxLifetime(d time.Duration) {
	db.SetConnMaxLifetime(d)
}
