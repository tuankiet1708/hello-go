package function

import "fmt"

func RunFunc() {
	i := 5
	j := 10
	fmt.Println(i, " - ", j) // 5 - 10
	i, j = swap(i, j)
	fmt.Println(i, " - ", j) // 10 - 5
}

func swap(i1, i2 int) (int, int) {
	return i2, i1
}
