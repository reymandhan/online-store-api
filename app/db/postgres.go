package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db *gorm.DB
)

func Init(host, port, username, password, dbName, sslMode string) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		username, password, host, port, dbName, sslMode)), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
				Colorful: true,
			},
		),
	})

	if err != nil {
		panic("Enable to connect to database: " + err.Error())
	}

	Db = db
}
