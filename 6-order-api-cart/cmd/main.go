package main

import (
	"6-project/configs"
	"6-project/internal/auth"
	"6-project/internal/order"
	"6-project/internal/product"
	"6-project/pkg/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)
	mux := mux.NewRouter()

	//REPOSITORY
	authRepository := auth.NewAuthRepo(db)
	productRepository := product.NewRepository(db)
	orderRepository := order.NewOrderRepository(db)

	//SERVICE
	serviceAuth := auth.NewServiceAuthByPhone(config, authRepository)
	serviceOrder := order.NewServiceOrder(&order.OrderServiceDesp{
		OrderRepository:   orderRepository,
		AuthByPhoneRepo:   authRepository,
		ProductRepository: productRepository,
	})

	//HANDLER
	auth.NewHandlerAuthByPhone(mux, &auth.AuthByPhoneHandlerDesp{
		ServiceAuthByPhone: serviceAuth,
		Config:             config,
	})
	product.NewProductHandler(mux, &product.ProductHandlerDesp{
		ProductRepository: productRepository,
	})

	order.NewHandlerOrder(mux, &order.OrderHandlerDesp{
		Config: config,
		OrderService:    serviceOrder,
	})

	//MIDDLEWARE
	// stack := middleware.Chain(
	// 	middleware.Auth,
	// )
	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	fmt.Println("Сервер запущен на порту 8081")
	log.Fatal(server.ListenAndServe())
}
