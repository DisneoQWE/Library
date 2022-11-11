package database

import (
	"RestApiLibrary/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToPostgresDB(c *config.Config) (*sqlx.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName)
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
