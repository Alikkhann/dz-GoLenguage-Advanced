package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func randHeaders(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(6)
	w.Write([]byte(strconv.Itoa(res)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/randHeader", randHeaders)
	server := http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}