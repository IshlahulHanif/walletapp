package database

import "github.com/jmoiron/sqlx"

type (
	ConfigDatabase struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
	}

	Module struct {
		dbConn *sqlx.DB
	}
)
