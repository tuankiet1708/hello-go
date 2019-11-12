package slicemap

import "fmt"

func RunSlice() {
	var a1 []int
	var a2 []int = []int{1, 2, 3}
	a3 := make([]int, 3)
	a4 := []int{1: 1, 3: 2}
	a5 := make([]int, 3, 5)
	fmt.Println(a1) // []
	fmt.Println(a2) // [1 2 3]
	fmt.Println(a3) // [0 0 0]
	fmt.Println(a4) // [0 1 0 2]
	fmt.Println(a5) // [0 0 0]
	a1 = append(a1, 1)
	a3 = append(a3, a2...)
	fmt.Println(a1)               // [1]
	fmt.Println(a3)               // [0 0 0 1 2 3]
	fmt.Println(copy(a4, a5), a4) // 3 [0 0 0 2]
}

func RunSliceAdvanced() {
	var a []int = []int{1, 2, 3, 4, 5}
	a1 := a[1:]
	fmt.Println(a1) // [2 3 4 5]
	a2 := a[:len(a)-1]
	fmt.Println(a2) // [1 2 3 4]
	a3 := a[:2]
	a3 = append(a3, a[3:]...)
	fmt.Println(a3) // [1 2 4 5]
	fmt.Println(a1) // [2 4 5 5]
	fmt.Println(a2) // [1 2 4 5]
	fmt.Println(a)  // [1 2 4 5 5]
}
