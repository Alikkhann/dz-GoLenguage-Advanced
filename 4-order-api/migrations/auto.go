package main

import (
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"4-project/internal/product"
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
	db.AutoMigrate(&product.Product{})
}