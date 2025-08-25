package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func NewGormDB() *GormDB {
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

	gormConfig := &gorm.Config{}

	db, err := gorm.Open(sqlite.Open(dbPath), gormConfig)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB:", err)
	}

	migrationsDir := filepath.Join(homeDir, "internal/database/migrations")
	if err := runMigrations(sqlDB, migrationsDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database connection established at:", dbPath)

	return &GormDB{
		DB: db,
	}
}

func (g *GormDB) Close() error {
	sqlDB, err := g.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
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

	goose.SetDialect("sqlite3")

	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
