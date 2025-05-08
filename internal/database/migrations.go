package database

import (
	"database/sql"
	"log"
	"os"
)

func Migrate(dbPath string) {
	// If no DB
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		createNewDB(dbPath)
	}
	// Open the file
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	gameTable(db)

	if err := db.Close(); err != nil {
		log.Fatalf("Failed to close DB file: %v", err)
	}
}

func createNewDB(dbPath string) {
	file, err := os.Create(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close new DB file: %v", err)
	}
	log.Println("Database file created at:", dbPath)
}

func gameTable(db *sql.DB) {
	var err error
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
