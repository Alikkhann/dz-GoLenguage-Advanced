package cmd

import (
	"4-project/configs"
	"4-project/internal/product"
	pkg "4-project/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	db := pkg.NewDb(config)
	db.AutoMigrate(&product.Product{})
}