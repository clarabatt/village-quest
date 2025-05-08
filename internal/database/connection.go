package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type QueryResult struct {
	Rows *sql.Rows
}

type DBAdapter interface {
	Query(statement string, params ...interface{}) (QueryResult, error)
	Exec(statement string, params ...interface{}) (sql.Result, error)
	Close() error
}

type SQLiteDB struct {
	connection *sql.DB
}

func NewSqliteAdapter() *SQLiteDB {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to get user home directory:", err)
	}

	appDir := filepath.Join(homeDir, ".village_quest")
	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		log.Fatal("Failed to create application directory:", err)
	}

	dbPath := filepath.Join(appDir, "village_quest.db")
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

func (s *SQLiteDB) Exec(statement string, params ...interface{}) (sql.Result, error) {
	result, err := s.connection.Exec(statement, params...)
	if err != nil {
		log.Printf("Error executing statement: %s, params: %v, error: %s", statement, params, err)
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
