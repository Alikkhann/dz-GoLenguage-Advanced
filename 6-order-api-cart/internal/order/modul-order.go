package order

import (
	"6-project/internal/product"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId uint
	Products []product.Product `gorm:"many2many:order_items;"`
	Status string
}

func NewOrder(userId uint, product []product.Product) *Order {
	return &Order{
		UserId: userId,
		Products: product,
		Status: "pending",
	}
}
