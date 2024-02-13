package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/eonianmonk/go-blog/assets"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	dbtype = "postgres"

	dbenv = "RPC_DB"
)

func getDb() (*sql.DB, error) {
	url := os.Getenv(dbenv)
	if url == "" {
		return nil, fmt.Errorf("no db env found")
	}
	return sql.Open(dbtype, url)
}

func runMigrate() {
	db, err := getDb()
	if err != nil {
		panic(fmt.Sprintf("failed to get db: %s", err.Error()))
	}
	fsmigrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: assets.Migrations,
		Root:       "sql/migrations",
	}
	applied, err := migrate.Exec(db, dbtype, fsmigrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("failed to apply migrations: %s", err.Error()))
	}
	log.Printf("applied %d migrations", applied)
}
