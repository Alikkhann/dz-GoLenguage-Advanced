package main

import (
    "fmt"
    "go/project-3/internal/repository"
    "go/project-3/internal/verify"
    "net/http"
)

func main() {
  repo := repository.VerifyRepository{FilePath: "dataRequest.json"}
  verify := verify.VerifyHandler{Repo: &repo}
  mux := http.NewServeMux()
  mux.HandleFunc("POST /send", verify.PostHandl)
  mux.HandleFunc("GET /verify/{hash}", verify.GetHandl)
  server := http.Server{
    Addr: ":8081",
    Handler: mux,
  }
  fmt.Println("Сервер запущен на порту 8081")
  server.ListenAndServe()
}