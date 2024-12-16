package database

import (
	"database/sql"
	"log"

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
	dbConn, err := sql.Open("sqlite3", "./village_quest.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &SqliteAdapter{
		connection: dbConn,
	}
}

func (s *SqliteAdapter) Query(statement string, params ...interface{}) (*sql.Rows, error) {
	return s.connection.Query(statement, params...)
}

func (s *SqliteAdapter) Close() error {
	return s.connection.Close()
}
