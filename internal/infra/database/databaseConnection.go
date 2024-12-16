package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type DatabaseConnection interface {
	Query(statement string, params ...interface{}) (*sql.Rows, error)
	Close() error
}

type SqliteAdapter struct {
	connection *sql.DB
}

func NewSqliteAdapter() *SqliteAdapter {
	Migrate()
	db, err := sql.Open("sqlite3", "./village_quest.db")

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &SqliteAdapter{
		connection: db,
	}
}

func (s *SqliteAdapter) Query(statement string, params ...interface{}) (*sql.Rows, error) {
	return s.connection.Query(statement, params...)
}

func (s *SqliteAdapter) Close() error {
	return s.connection.Close()
}

func Migrate() {
	dbPath := "./village_quest.db"
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
        id UUID PRIMARY KEY,
        number INT NOT NULL,
        max_days_played INT NOT NULL,
        players_name VARCHAR(255) NOT NULL
    );
    `

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'game' created or already exists.")
}
