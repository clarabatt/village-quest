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
	Err  error
}

type DBAdapter interface {
	Query(statement string, params ...interface{}) QueryResult
	Exec(statement string, params ...interface{}) (sql.Result, error)
	Close() error
}

type SQLiteDB struct {
	connection *sql.DB
}

func NewSqliteAdapter() *SQLiteDB {
	Migrate()
	dbPath := filepath.Join(dir(), "village_quest.db")
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
	log.Printf("Statement executed: %s, params: %v", statement, params)
	return result, nil
}

func (s *SQLiteDB) Query(statement string, params ...interface{}) QueryResult {
	rows, err := s.connection.Query(statement, params...)
	log.Println("Query executed:", statement)
	log.Println("Params:", params)
	log.Println("Rows:", rows)
	if err != nil {
		log.Panicf("Error executing query: %s", err)
	}
	return QueryResult{rows, err}
}

func (s *SQLiteDB) Close() error {
	return s.connection.Close()
}

func dir() string {
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(executablePath)
}

func Migrate() {
	dbPath := filepath.Join(dir(), "village_quest.db")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		log.Println("Database file created:", dbPath)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS game (
        id TEXT PRIMARY KEY,
        number INTEGER NOT NULL,
        max_days_played INTEGER NOT NULL,
        players_name TEXT NOT NULL
    );
    `

	_, err = db.Exec(createTableQuery)

	if err != nil {
		log.Fatal(err)
		log.Println("Error creating 'game' table")
	} else {
		log.Println("Table 'game' ok")
	}
}
