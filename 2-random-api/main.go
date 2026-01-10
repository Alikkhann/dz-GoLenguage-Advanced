package main

import (
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func randHeaders(w http.ResponseWriter, r *http.Request) {
	res := rand.IntN(6)
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