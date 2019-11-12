package main

import (
	"fmt"
	"os"
	"strconv"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	i, _ := strconv.Atoi(os.Args[1])
	fmt.Println(fact(i))
}
