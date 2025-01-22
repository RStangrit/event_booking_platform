package database

import (
	"database/sql"
	"main/migrations"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	if DB, err = sql.Open("sqlite3", "api.db"); err != nil {
		panic("could not connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	migrations.CreateTables(DB)
}
