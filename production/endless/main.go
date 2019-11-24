package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fvbock/endless"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/hello", handler).Methods("GET")

	err := endless.ListenAndServe(":4242", mux)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server ở port 4242 đã ngừng!")

	os.Exit(0)
}
