package main

import (
	"fmt"
	"net/http"
	"time"
)

func liteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chào mừng đến lập trình Go cho web!")
}

// func main() {
//     mux := http.NewServeMux()
//     // Ép kiểu liteHandler về http.HandleFunc
//     hdl := http.HandlerFunc(liteHandler)
//     mux.Handle("/welcome", hdl)
//     http.ListenAndServe(":8080", mux)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/welcome", liteHandler)
// 	http.ListenAndServe(":8080", mux)
// }

// func main() {
// 	http.HandleFunc("/welcome", liteHandler)
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	http.HandleFunc("/welcome", liteHandler)
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
