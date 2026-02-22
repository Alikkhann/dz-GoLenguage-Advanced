package main

import (
	"fmt"
	"log"
	"net/http"
	"4-project/pkg/db"
	"4-project/configs"
	"github.com/gorilla/mux"
	"4-project/internal/product"
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
		Handler: mux,
	}
	fmt.Println("Сервер запущен на порту 8081")
	log.Fatal(server.ListenAndServe())
}