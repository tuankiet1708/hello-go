package intstring

import (
	"fmt"
	"strconv"
)

func Run() {
	var i int = 1234
	s1 := fmt.Sprintf("%d", i)
	s2 := fmt.Sprintf("%x", i)
	s3 := strconv.Itoa(i)
	s4 := strconv.FormatInt(int64(i), 16)
	fmt.Println(s1) // "1234"
	fmt.Println(s2) // "4d2" là giá trị của 1234 trong hệ 16
	fmt.Println(s3) // "1234"
	fmt.Println(s4) // "4d2" là giá trị của 1234 trong hệ 16

	s := "1234"
	x1, _ := strconv.Atoi(s)
	x2, _ := strconv.ParseInt(s, 10, 32)
	s = "4d2"
	x3, _ := strconv.ParseInt(s, 16, 32)
	fmt.Println(x1) // 1234
	fmt.Println(x2) // 1234
	fmt.Println(x3) // 1234  là giá trị trong hệ 10, ứng với 4d2 trong hệ 16
}
