package product

import "gorm.io/gorm"

type Product struct {
  gorm.Model
  Name        string
  Description string
  // Images      pq.StringArray
}

func NewProduct() {
	
}