package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	_, err := fmt.Fprintln(f, "data")
	fmt.Println("error", err)
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
