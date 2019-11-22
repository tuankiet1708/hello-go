package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/pat"
)

var token string

func indexHandler(w http.ResponseWriter, req *http.Request) {
	h := md5.New()
	curtime := time.Now().Unix()
	io.WriteString(h, strconv.FormatInt(curtime, 10))
	token = fmt.Sprintf("%x", h.Sum(nil))
	t, _ := template.ParseFiles("form.gtpl")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, token)
}

func validateHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	postToken := req.PostForm.Get("token")
	if postToken != "" {
		// Thực hiện so sánh với token đã tạo trước đó
		if postToken == token {
			fmt.Println("Token matched")
		} else {
			fmt.Println("Token not existed")
		}
	} else {
		// Token không tồn tại, lỗi!
		fmt.Println("Token not existed")
	}
}

func main() {
	r := pat.New()
	r.Get("/", indexHandler)
	r.Post("/validate", validateHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
