package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// import "encoding/base64"
	rawstr := "Chào thế giới Go!"
	// Base64 chuẩn
	stdBase64 := base64.StdEncoding.EncodeToString([]byte(rawstr))
	fmt.Println(stdBase64)
	// decode
	deStd, err := base64.StdEncoding.DecodeString(stdBase64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(deStd))
	// Base64 phù hợp URL, Raw sẽ loại bỏ padding
	urlBase64 := base64.RawURLEncoding.EncodeToString([]byte(rawstr))
	fmt.Println(urlBase64)
	// decode
	deURL, err := base64.RawURLEncoding.DecodeString(urlBase64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(deURL))
}
