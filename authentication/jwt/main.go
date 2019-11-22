package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type JWT struct {
	Token string `json:"token"`
}

var secret []byte = []byte("secrethmac256jwt")

func jwtHandler(w http.ResponseWriter, r *http.Request) {
	// create JWT
	claims := MyClaims{
		"root",
		"admin",
		jwt.StandardClaims{
			Issuer:    "pilo",
			ExpiresAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token.Header)
	fmt.Println(token.Claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatal("Lỗi tạo JWT:", err.Error())
	}
	fmt.Println(tokenString)

	// send JWT
	tokenstr := JWT{tokenString}
	json, err := json.Marshal(tokenstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

type JWTValidate struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	ExpiresAt int64  `json:"exp"`
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := string(r.URL.Query()["token"][0])
	fmt.Println(r.URL.Query(), tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	var response JWTValidate
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		fmt.Printf("%v %v %v", claims.Username, claims.Role, claims.StandardClaims.ExpiresAt)
		response = JWTValidate{tokenString, claims.Username, claims.Role, claims.StandardClaims.ExpiresAt}
	} else {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(response)
	w.Write(json)
}

func main() {
	http.HandleFunc("/jwt", jwtHandler)
	http.HandleFunc("/validate", validateHandler)
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
