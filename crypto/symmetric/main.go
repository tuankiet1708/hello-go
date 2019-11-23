package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func encode(plaintext []byte, c cipher.Block) []byte {
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}

func decode(ciphertext []byte, c cipher.Block) []byte {
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, 65536)
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)

	return plaintextCopy
}

func main() {
	plaintext := []byte("I love You")
	keytext := "ABBAabbaBAABbaab"

	c, err := aes.NewCipher([]byte(keytext))
	if err != nil {
		log.Fatalf("Lỗi %s", err.Error())
	}

	ciphertext := encode(plaintext, c)
	fmt.Printf("Encode: %s => %x\n", plaintext, ciphertext)
	plaintext = decode(ciphertext, c)
	fmt.Printf("Decode: %x => %s\n", ciphertext, plaintext)
}
