package database

import "database/sql"

var DB *sql.DB

func InitDB() {
	var err error

	if DB, err = sql.Open("sqlite3", "api.db"); err != nil {
		panic("could not connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// createTables()
}
