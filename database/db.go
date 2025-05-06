package database

import (
	"fmt"
	"test1/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=test port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
	}
	db.AutoMigrate(&models.Product{}, &models.Measure{})

	DB = db
}
