package main

import (
	"5-project/configs"
	"5-project/internal/authByPhone"
	"5-project/pkg/db"
	"5-project/pkg/middleware"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	//конфигурации
	config := configs.LoadConfig()

	//база данных
	db := db.NewDb(config)

	//REPOSITORY
	repo := authbyphone.NewAuthRepo(db)

	//роутер
	mux := mux.NewRouter()
	//SERVICE
	service := authbyphone.NewServiceAuthByPhone(config, repo)

	//HANDLER
	authbyphone.NewHandlerAuthByPhone(mux, &authbyphone.AuthByPhoneHandlerDesp{
		ServiceAuthByPhone: service,
	})

	//MIDDLEWARE
	stack := middleware.Chain(
		middleware.Auth,
	)


	server := http.Server{
		Addr:    ":8081",
		Handler: stack(mux),
	}

	fmt.Println("Сервер запущен на порту 8081")
	log.Fatal(server.ListenAndServe())
}
