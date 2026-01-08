package database

import (
	"database/sql"
	"embed"
	"log"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func SetupDatabase() *sql.DB {
	db, err := sql.Open("sqlite", "./swiss.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	err = goose.SetDialect("sqlite3")
	if err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	goose.SetBaseFS(embedMigrations)

	err = goose.Up(db, "migrations")
	if err != nil {
		log.Fatalf("goose up failed: %v", err)
	}

	fmt.Printf("Migrations completed successfully!")
	return db
}
