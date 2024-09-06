package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/mattn/go-sqlite3"

	"github.com/tealok-tech/tealok/database/migrations"
	"github.com/tealok-tech/tealok/database/sqlc/procedures"
)

const DB_FILE = "/var/lib/tealok/database.sqlite"

func Connect() (*sql.DB, error) {
	log.Println("Connecting to database")
	// Connect up and run any migrations that are outstanding
	s := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})
	d, err := bindata.WithInstance(s)
	m, err := migrate.NewWithSourceInstance("go-bindata", d, "sqlite3://"+DB_FILE)
	if err != nil {
		return nil, err
	}
	m.Up()

	version, dirty, err := m.Version()
	if err != nil {
		return nil, err
	}
	log.Println("Database is now at version", version, dirty)
	m.Close()

	// Geta new connection where we can make queries
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return nil, err
	}

	var sqlversion string
	row := db.QueryRow("select sqlite_version()")
	err = row.Scan(&sqlversion)
	if err != nil {
		return nil, err
	}
	log.Println("SQLite version:", sqlversion)
	return db, nil
}

var ddl string

func AddContainer(db *sql.DB, name string) error {
	if db == nil {
		return errors.New("Nil database connection")
	}
	ctx := context.Background()
	queries := procedures.New(db)
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Format it as a string suitable for SQLite (ISO 8601 format)
	formattedTime := currentTime.Format(time.RFC3339)

	// Print the formatted time
	fmt.Println("Formatted Time:", formattedTime)

	// This formattedTime can now be inserted into an SQLite database

	container, err := queries.CreateContainer(ctx, procedures.CreateContainerParams{
		Name:      name,
		CreatedAt: formattedTime,
	})
	if err != nil {
		return err
	}
	log.Println("Added container", container)
	return nil
}
