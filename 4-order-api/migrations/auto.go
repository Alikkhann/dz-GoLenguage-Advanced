package migrations

import (
	"4-project/configs"
	"4-project/internal/product"
	"4-project/pkg"
)

func main() {
	config := configs.Config{}
	db := pkg.NewDb(&config)
	db.AutoMigrate(product.Product{})
}