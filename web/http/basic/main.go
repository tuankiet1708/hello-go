package main

import (
	"fmt"
	"net/http"
)

type liteHandler struct {
	message string
}

func (m *liteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

// func main() {
// 	hdl := &liteHandler{"Chào mừng đến lập trình Go cho web!"}
// 	http.ListenAndServe(":8080", hdl)
// }
func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("html"))
	mux.Handle("/", fs)
	hdl := &liteHandler{"Chào mừng đến lập trình Go cho web!"}
	mux.Handle("/welcome", hdl)
	http.ListenAndServe(":8080", mux)
}
