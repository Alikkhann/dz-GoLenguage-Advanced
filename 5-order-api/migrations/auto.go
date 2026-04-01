package main

import (
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"5-project/internal/authByPhone"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), (&gorm.Config{}))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&authbyphone.AuthByPhone{})
}