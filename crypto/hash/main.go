package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func runSha256() {
	// import "crypto/sha256"
	h := sha256.New()
	pw := "asdfghjkl"
	io.WriteString(h, pw)
	pwsha256 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(pwsha256)
	// 5c80565db6f29da0b01aa12522c37b32f121cbe47a861ef7f006cb22922dffa1
}

func runSha1() {
	// import "crypto/sha1"
	h := sha1.New()
	pw := "qwertyuiop"
	io.WriteString(h, pw)
	pwsha1 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(pwsha1) // 5fa339bbbb1eeaced3b52e54f44576aaf0d77d96
}

func runMd5() {
	// import "crypto/md5"
	h := md5.New()
	pw := "qwertyuiop"
	io.WriteString(h, pw)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(pwmd5) // c44a471bd78cc6c2fea32b9fe028d30a
}

func runMd5WithSalt() {
	// import "crypto/md5"
	h := md5.New()
	username := "abc"
	password := "zxcvbnm"
	io.WriteString(h, password)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(pwmd5) //Kết quả: 02c75fb22c75b23dc963c7eb91a062cc

	// Bổ sung 2 thành phần:
	salt1 := "@#$%"
	salt2 := "^&*()"

	//Xóa nội dung chuỗi băm
	h.Reset()

	// salt1 + username + salt2 + PWMD5
	io.WriteString(h, salt1)
	io.WriteString(h, username)
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)
	finalpw := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(finalpw) //Kết quả: a58ac84c53eb80cff5169cdc0ca1a32f
}

func main() {
	fmt.Println("SHA-256")
	runSha256()
	fmt.Println()
	fmt.Println("SHA-1")
	runSha1()
	fmt.Println()
	fmt.Println("MD5")
	runMd5()
	fmt.Println()
	fmt.Println("MD5 with salt")
	runMd5WithSalt()
	fmt.Println()
}
