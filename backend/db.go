package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title  string `json:"Title" binding:"required"`
	Body   string `json:"Body" binding:"required"`
	UserID int
	User   User `json:"owner" binding:"required"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	is_admin bool   `json:"is_admin"`
}

func initDB() (*gorm.DB, error) {

	// for docker
	dsn := "host=db user=postgres password=postgres dbname=docker port=5432 sslmode=disable"
	// for localhost
	// dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&Note{})
	db.AutoMigrate(&User{})
	return db, err
}
