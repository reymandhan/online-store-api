package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	Global *Config
)

// DatabaseConfig store database related configurations
type DatabaseConfig struct {
	Host     string `envconfig:"OS_API_DB_HOST" default:"db"`
	Port     string `envconfig:"OS_API_DB_PORT" default:"5432"`
	Username string `envconfig:"OS_API_DB_USERNAME" default:"user"`
	Password string `envconfig:"OS_API_DB_PASSWORD" default:"password"`
	Name     string `envconfig:"OS_API_DB_NAME" default:"db"`
	SSLMode  string `envconfig:"OS_API_DB_SSL_MODE" default:"disable"`
}

// Application related config
type Config struct {
	Environment string `envconfig:"OS_API_ENV" default:"DEV"`
	Port        string `envconfig:"OS_API_PORT" default:":8080"`
	APIPrefix   string `envconfig:"OS_API_API_PREFIX" default:"/api/v1/"`
	Database    *DatabaseConfig
}

func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	conf := new(Config)
	envconfig.MustProcess("os_api", conf)
	database := new(DatabaseConfig)
	envconfig.MustProcess("os_api", database)

	conf.Database = database
	Global = conf
}
