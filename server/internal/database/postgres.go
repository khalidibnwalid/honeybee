package database

import (
	"fmt"
	"net/url"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)


type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
	TimeZone string
}

// postgres URI: postgres://<username>:<password>@<host>:<port>/<dbname>?sslmode=[disable]
func ParsePostgresURI(uri string) (*PostgresConfig, error) {
	parsedURL, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URI: %w", err)
	}

	if parsedURL.Scheme != "postgres" {
		return nil, fmt.Errorf("invalid scheme: %s, expected 'postgres://'", parsedURL.Scheme)
	}

	username := ""
	password := ""
	if parsedURL.User != nil {
		username = parsedURL.User.Username()
		password, _ = parsedURL.User.Password()
	}

	host := parsedURL.Hostname()
	port := parsedURL.Port()
	if port == "" {
		port = "5432" // Default PostgreSQL port
	}

	database := strings.TrimPrefix(parsedURL.Path, "/")

	sslMode := "disable"
	timeZone := "UTC"
	if parsedURL.Query().Get("sslmode") != "" {
		sslMode = parsedURL.Query().Get("sslmode")
	}
	if parsedURL.Query().Get("TimeZone") != "" {
		timeZone = parsedURL.Query().Get("TimeZone")
	}

	return &PostgresConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
		SSLMode:  sslMode,
		TimeZone: timeZone,
	}, nil
}

