package database

import (
	"fmt"
	"log"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/tealok-tech/tealok/database/migrations"
)

func Connect() {
	fmt.Println("Connecting to database")
	s := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})
	d, err := bindata.WithInstance(s)
	m, err := migrate.NewWithSourceInstance("go-bindata", d, "sqlite:///var/lib/tealok/database.sqlite")
	if err != nil {
		log.Fatal(err)
		return
	}
	m.Up()
}
