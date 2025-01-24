package database

import (
	"database/sql" //for sqlite
	"fmt"
	"log"
	"main/migrationsPostgres"
	"main/migrationsSqlite" //lauch sqlite migrations
	"main/pkg/util"         //for reading .env DSN

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *sql.DB
	ORMDB *gorm.DB
)

// InitDB initializes the database based on the given dbOperator
func InitDB(dbOperator string) {
	switch dbOperator {
	case "sqlite":
		if err := initSQLiteDB(); err != nil {
			log.Fatalf("Failed to initialize SQLite: %v", err)
		}
		log.Println("SQLite database initialized successfully.")
	case "postgres":
		if err := initPostgresDB(); err != nil {
			log.Fatalf("Failed to initialize PostgreSQL: %v", err)
		}
		log.Println("PostgreSQL database initialized successfully.")
	default:
		panic("could not connect to any database")
	}
}

// initSQLiteDB initializes SQLite database
func initSQLiteDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		return fmt.Errorf("could not connect to the SQLite database: %w", err)
	}

	// Set connection pool parameters
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Run migrations
	if err := migrationsSqlite.CreateTablesSql(DB); err != nil {
		return fmt.Errorf("failed to run SQLite migrations: %w", err)
	}

	return nil
}

// initPostgresDB initializes PostgreSQL database
func initPostgresDB() error {
	var err error
	dsn := util.GetEnvVariable("DSN")
	if dsn == "" {
		return fmt.Errorf("DSN environment variable is not set")
	}

	ORMDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not connect to the PostgreSQL database: %w", err)
	}

	// Run migrations
	if err := migrationsPostgres.CreateTablesPostgres(ORMDB); err != nil {
		return fmt.Errorf("failed to run PostgreSQL migrations: %w", err)
	}

	return nil
}
