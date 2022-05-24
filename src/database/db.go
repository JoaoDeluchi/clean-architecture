package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&User{})

	return db
}

func New(db *gorm.DB) handler {
	return handler{db}
}
