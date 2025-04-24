package database

import (
	"database/sql"
	"log"
	"os"
)

func Migrate(dbPath string) {
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

	gameTable(db)
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
