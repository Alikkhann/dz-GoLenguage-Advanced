package main

import (
	"5-project/configs"
	"5-project/internal/authByPhone"
	"5-project/pkg/db"
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

	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	fmt.Println("Сервер запущен на порту 8081")
	log.Fatal(server.ListenAndServe())
}
