package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Install pgvector extension
	_, err = db.Exec("CREATE EXTENSION IF NOT EXISTS vector;")
	if err != nil {
		log.Printf("Warning: Could not create vector extension: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Printf("Error during migration: %v", err)
		log.Println("Attempting to fix...")

		// Get current version
		version, dirty, _ := m.Version()
		if dirty {
			log.Printf("Forcing version %d", version)
			err = m.Force(int(version))
			if err != nil {
				log.Printf("Could not force version: %v", err)
			}
		}

		// Drop migrations table and try again
		log.Println("Dropping schema_migrations table...")
		_, err = db.Exec("DROP TABLE IF EXISTS schema_migrations;")
		if err != nil {
			log.Fatalf("Could not drop schema_migrations table: %v", err)
		}

		log.Println("Retrying migration...")
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed after reset: %v", err)
		}
	}

	log.Println("Migration completed successfully")
}
