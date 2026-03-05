package main

import (
	"4-project/configs"
	"4-project/internal/product"
	"4-project/pkg/db"
	"4-project/pkg/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	mux := mux.NewRouter()
	
	//repo
	repo := product.NewRepository(db)

	//handler product
	product.NewProductHandler(mux, product.ProductHandlerDesp{
		ProductRepository: repo,
	})
	
	server := http.Server{
		Addr: ":8081",
		Handler: middleware.Logger(mux),
	}
	fmt.Println("Сервер запущен на порту 8081")
	log.Fatal(server.ListenAndServe())
}