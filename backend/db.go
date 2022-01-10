package main

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title  string `json:"Title" binding:"required"`
	Body   string `json:"Body" binding:"required"`
	UserID int
}

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Is_admin bool   `json:"is_admin"`
}

func initDB() gorm.DB {

	// for docker
	// dsn := "host=db user=postgres password=postgres dbname=docker port=5432 sslmode=disable"
	// for localhost
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") +
		" password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&Note{})
	db.AutoMigrate(&User{})
	return *db
}
