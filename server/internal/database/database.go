package database

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const envPgURI = "POSTGRES_URI"

type Database struct {
	Client *gorm.DB
}

// if uri is not provided, it will look for PG_URI in the environment variables
func NewClient(uri ...string) (*Database, error) {
	var _uri string
	if len(uri) > 0 {
		_uri = uri[0]
	} else {
		_uri = os.Getenv(envPgURI)
	}

	if _uri == "" {
		return nil, fmt.Errorf("no URI provided and %v environment variable is not set", envPgURI)
	}

	pg, err := ParsePostgresURI(_uri)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Postgres URI: %w", err)
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", pg.Host, pg.Username, pg.Password, pg.Database, pg.Port, pg.SSLMode, pg.TimeZone)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Database{
		Client: db,
	}, nil
}

func (db *Database) Ping() error {
	return db.Client.Exec("SELECT 1").Error
}