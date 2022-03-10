package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"test-api/pkg/models"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5435/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
