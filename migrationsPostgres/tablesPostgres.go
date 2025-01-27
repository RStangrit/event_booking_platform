package migrationsPostgres

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         int
	Name       string `binding:"required"`
	Email      string `binding:"required"`
	Password   string `binding:"required"`
	Role       string `binding:"required"`
	Created_at time.Time
	Updated_at *time.Time
}

type Event struct {
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

func CreateTablesPostgres(GORM *gorm.DB) error {
	err := createTable(GORM)
	return err
}

func createTable(GORM *gorm.DB) error {
	err := GORM.AutoMigrate(&User{}, &Event{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
		return err
	}
	log.Println("Migration completed successfully")
	return nil
}
