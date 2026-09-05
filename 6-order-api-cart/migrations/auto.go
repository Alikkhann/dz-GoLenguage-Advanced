package main

import (
	"6-project/internal/auth"
	"6-project/internal/order"
	"6-project/internal/product"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db.AutoMigrate(&auth.AuthByPhone{}, &product.Product{}, &order.Order{})
}