package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

type Seed struct {
	db *sql.DB
}

func InitSeed(host, port, username, password, dbName, sslMode string) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, dbName, sslMode)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	Execute(db)
}

func Execute(db *sql.DB) {
	q, err := ioutil.ReadFile(GetSourcePath() + "/seed/init_data.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}

func GetSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
