package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
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

	// Check if the database file exists, if not, create a new one
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		createNewDB(dbPath)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	goose.SetDialect("sqlite3")

	migrationsDir := filepath.Join(homeDir, "internal/database/migrations")
	if err := runMigrations(db, migrationsDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
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

func createNewDB(dbPath string) {
	file, err := os.Create(dbPath)
	if err != nil {
		log.Fatal("Failed to create database file:", err)
	}
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close new DB file: %v", err)
	}
	log.Println("Database file created at:", dbPath)
}

func runMigrations(db *sql.DB, migrationsDir string) error {
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(migrationsDir, 0755); err != nil {
			return fmt.Errorf("failed to create migrations directory: %w", err)
		}
		log.Printf("Created migrations directory: %s", migrationsDir)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
