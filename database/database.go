package database

import (
	"database/sql"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/mattn/go-sqlite3"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/tealok-tech/tealok/database/migrations"
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
	m, err := migrate.NewWithSourceInstance("go-bindata", d, "sqlite://"+DB_FILE)
	if err != nil {
		return nil, err
	}
	m.Up()
	m.Close()

	// Geta new connection where we can make queries
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return nil, err
	}

	var version string
	row := db.QueryRow("select sqlite_version()")
	err = row.Scan(&version)
	if err != nil {
		return nil, err
	}
	log.Println("SQLite version:", version)
	return db, nil
}
