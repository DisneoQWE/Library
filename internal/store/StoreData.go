package store

import (
	"github.com/jmoiron/sqlx"
)

type DBConnection struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *DBConnection {
	return &DBConnection{
		db: db,
	}
}
