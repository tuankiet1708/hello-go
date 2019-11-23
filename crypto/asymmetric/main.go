package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
)

var aMessage []byte = []byte("Anh yêu em!")
var bPrivateKey *rsa.PrivateKey
var bPublicKey *rsa.PublicKey
var err error

func encodedecode() {
	// Tạo khoá
	bPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	bPublicKey = &bPrivateKey.PublicKey

	fmt.Println("Khóa bí mật: ", bPrivateKey)
	fmt.Println("Khóa công khai: ", bPublicKey)

	// Mã hoá
	label := []byte("")

	var ciphertext []byte
	ciphertext, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, bPublicKey, aMessage, label)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	fmt.Printf("Mã hóa OAEP [%s] thành \n[%x]\n", string(aMessage), ciphertext)

	// Giải mã
	var messageFromA []byte
	messageFromA, err = rsa.DecryptOAEP(sha256.New(), rand.Reader, bPrivateKey, ciphertext, label)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	fmt.Printf("Giải mã OAEP [%x] được \n[%s]\n", ciphertext, messageFromA)
}

func signature() {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	hashString := sha256.Sum256(aMessage)
	var aSignature []byte
	aSignature, err = rsa.SignPSS(rand.Reader, bPrivateKey, crypto.SHA256, hashString[:], &opts)
	if err != nil {
		log.Fatalf("Lỗi: %s", err.Error())
	}
	fmt.Printf("Chữ ký: %x\n", aSignature)

	err = rsa.VerifyPSS(bPublicKey, crypto.SHA256, hashString[:], aSignature, &opts)
	if err != nil {
		log.Fatal("Lỗi xác thực chữ ký")
	} else {
		fmt.Println("Chữ ký được xác thực của A!")
	}
}

func main() {
	encodedecode()
	signature()
}
