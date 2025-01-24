package migrationsPostgres

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id         int
	Name       string `binding:"required"`
	Email      string `binding:"required"`
	Password   string `binding:"required"`
	Role       string `binding:"required"`
	Created_at time.Time
	Updated_at *time.Time
}

type EventModel struct {
	ID          int
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Date        string `binding:"required"`
	Location    string `binding:"required"`
	Capacity    int    `binding:"required"`
	Price       *float64
	Created_by  int `binding:"required"`
	Created_at  time.Time
	Updated_at  *time.Time
}

var tables = []interface{}{
	&UserModel{},
	&EventModel{},
}

func CreateTablesPostgres(ORMDB *gorm.DB) error {
	err := createTable(tables, ORMDB)
	return err
}

func createTable(tables []interface{}, ORMDB *gorm.DB) error {
	err := ORMDB.AutoMigrate(&UserModel{}, &EventModel{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
		return err
	}
	log.Println("Migration completed successfully")
	return nil
}
