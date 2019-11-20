package main

import (
	"fmt"
	"os"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// var f func(int) int
// f = func (n int) int {
//     ...
//     f(x)
//     ...
// }

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func makeDirs() {
	if len(os.Args) < 2 {
		fmt.Println("Vui lòng nhập tên các thư mục cần tạo: dir thưmục1 ... thưmụcn")
		return
	}

	var rmdirs []func()
	for _, dir := range os.Args[1:] {
		d := dir // Quan trọng!
		os.MkdirAll(d, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(d)
		})
	}
	fmt.Print("Các thư mục đã được tạo. Nhấn phím bất kỳ để xóa chúng!")
	var c string
	fmt.Scanf("%s", &c)
	for index, rmdir := range rmdirs {
		fmt.Println("index", index)
		rmdir() // Xóa thư mục
	}
}

func main() {
	// f := squares()
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())

	// makeDirs()

	var list []int = []int{1, 2, 3, 4, 5}
	res := sum(list...)
	fmt.Println("Sum = ", res)
}
