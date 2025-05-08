package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	connection *sql.DB
}

type QueryResult struct {
	Rows *sql.Rows
}

type DBAdapter interface {
	Query(statement string, params ...interface{}) (QueryResult, error)
	Execute(statement string, params ...interface{}) (sql.Result, error)
	Close() error
}

func NewSqliteAdapter() *SQLiteDB {
	homeDir, err := os.Getwd()
	fmt.Printf("Home Dir: %v\n", homeDir)
	if err != nil {
		log.Fatal("Failed to get project's home directory:", err)
	}

	dbPath := filepath.Join(homeDir, "data.db")
	log.Println("Database path:", dbPath)

	Migrate(dbPath)

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established at:", dbPath)

	return &SQLiteDB{
		connection: db,
	}
}

func (s *SQLiteDB) Execute(statement string, params ...interface{}) (sql.Result, error) {
	result, err := s.connection.Exec(statement, params...)
	if err != nil {
		log.Panicf("Error executing statement: %s, params: %v, error: %s", statement, params, err)
		return nil, err
	}
	return result, err
}

func (s *SQLiteDB) Query(statement string, params ...interface{}) (QueryResult, error) {
	rows, err := s.connection.Query(statement, params...)
	if err != nil {
		log.Panicf("Error executing query: %s", err)
	}
	return QueryResult{rows}, err
}

func (s *SQLiteDB) Close() error {
	return s.connection.Close()
}
