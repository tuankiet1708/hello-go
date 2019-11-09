package arraystruct

import (
	"fmt"
)

func RunArray() {
	var a1 [2]int
	var a2 [3]int = [3]int{1, 2, 3}
	a3 := [3]int{2: 4, 1: 5}
	fmt.Println(a1)      // [0 0]
	fmt.Println(a2)      // [1 2 3]
	fmt.Println(a3)      // [4 5 0]
	fmt.Println(len(a1)) // 2
	fmt.Println(len(a2)) // 3
	fmt.Println(len(a3)) // 3
	fmt.Println(a2[2])   // 3
	fmt.Println(a3[2])   // 0
	// fmt.Println(a3[3])   // Lỗi khi biên dịch

	a4 := [...]int{1, 2, 3}
	// a4 = [...]int{1, 2, 3, 4, 5, 6, 7} // Lỗi khi biên dịch
	// a4[4] = 7                          // Lỗi khi biên dịch
	a4 = [3]int{}
	fmt.Println(a4)

	var a5 [3]int
	var a6 [3]int = [3]int{1, 2, 3}
	a7 := [3]int{1, 2}
	fmt.Println(a5 == a6) // false
	fmt.Println(a6 == a7) // false
	a7[2] = 3
	fmt.Println(a7 == a6) // true
	fmt.Println(a5 != a6) // true
	fmt.Println(a6 != a7) // false
}
