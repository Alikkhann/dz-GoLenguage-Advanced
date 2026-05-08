package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Name        string
  Description string
  Images      pq.StringArray `gorm:"type:text[]"`
}

func NewProduct(prodBody ProductCreateRequest) *Product{
  return &Product{
    Name: prodBody.Name,
    Description: prodBody.Description,
    Images: prodBody.Image,
  }
}


