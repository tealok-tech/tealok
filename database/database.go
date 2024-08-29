package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
	//_ "github.com/golang-migrate/migrate/v4"
	//_ "github.com/golang-migrate/migrate/v4/database/sqlite"
)

func Connect() {
	fmt.Println("Connecting to database")
	db, err := sql.Open("sqlite", "/var/lib/tealok/database.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	var version string
	row := db.QueryRow("select sqlite_version()")
	err = row.Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SQLite version:", version)
}
