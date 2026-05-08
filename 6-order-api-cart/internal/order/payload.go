package order

import "6-project/internal/product"

type OrderCreateRequest struct {
	ProductsID []uint `json:"products_id" validate:"required,min=1,dive,gt=0"`
}

type OrderResponse struct {
	ID uint `json:"id"`
	UserId uint `json:"userId"`
	Product []product.Product `json:"product"`
}

// required	Слайс не может быть nil
// min=1	Минимум 1 элемент в слайсе
// dive	"Ныряет" внутрь слайса — дальше правила применяются к каждому элементу
// gt=0	Каждый ID должен быть больше 0