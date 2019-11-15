package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

// func main() {
// 	i, _ := strconv.Atoi(os.Args[1])
// 	fmt.Println(fact(i))
// }

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Vui lòng thêm chuỗi và khóa: caesar <chuỗi> <khóa>")
		return
	}

	k, _ := strconv.Atoi(os.Args[2])
	caesar := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+rune(k))%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+rune(k))%26
		}
		return r
	}

	fmt.Println(strings.Map(caesar, os.Args[1]))

}
