package dbDriver

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

//const maxOpenDbConn = 20
//const maxIdleDbConn = 3
//const connMaxLifetime = 1 * time.Second

// ConnectSQL creates db pool for sqlite
func ConnectSQL(dbName string) (*DB, error) {
	db, err := NewDatabase(dbName)
	if err != nil {
		panic(err)
	}

	//db.SetMaxOpenConns(maxOpenDbConn)
	//db.SetMaxIdleConns(maxIdleDbConn)
	//db.SetConnMaxLifetime(connMaxLifetime)
	//db.SetConnMaxIdleTime(connMaxLifetime)

	dbConn.SQL = db

	return dbConn, nil
}

// testDB tries to ping the database
func testDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new database for the application
func NewDatabase(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("db/%s.db", dbName))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
