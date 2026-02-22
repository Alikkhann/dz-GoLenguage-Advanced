package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name string	`json:"name" validate: "required"`
	Description string `json:"description" validate: "required"`
	Image pq.StringArray `json:"images" validate: "required, dive, url"`
}

//required: Поле обязательно для заполнения.
//dive: Используется для вложенного уровня валидации, применяется для валидации каждого элемента в слайсе или массиве.
//url: Проверяет, что каждая строка в массиве является действительным URL (если библиотека поддерживает эту валидацию).