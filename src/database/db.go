package database

import (
	"log"

	"github.com/JoaoDeluchi/clean-architecture/src/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entity.User{})

	return db
}
