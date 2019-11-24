package main

import (
	"io"
	"log"
	"os"
)

var (
	lInfo    *log.Logger
	lWarning *log.Logger
	lError   *log.Logger
)

func Init(out io.Writer) {
	flag := log.LstdFlags | log.Lshortfile
	lInfo = log.New(out, "INFO: ", flag)
	lWarning = log.New(out, "WARNING: ", flag)
	lError = log.New(out, "ERROR: ", flag)
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	Init(f)
	lInfo.Println("Thông tin")
	lWarning.Println("Cảnh báo")
	lError.Println("Lỗi")
}
