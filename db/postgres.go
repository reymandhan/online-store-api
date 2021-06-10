package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	Db *sqlx.DB
)

func Init(host, port, username, password, dbName, sslMode string) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, username, password, dbName, sslMode))

	if err != nil {
		log.Print(err)
	}
	if err = db.Ping(); err != nil {
		log.Print(err)
	}

	Db = db
}

func Ping() bool {
	if err := Db.Ping(); err != nil {
		return false
	}
	return true
}

func Migrate(host, port, username, password, dbName, sslMode string) error {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, dbName, sslMode)

	m, err := migrate.New("file://db/migrations", connString)
	if err != nil {
		log.Print(err)
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no change"):
		return nil
	}

	return nil
}
