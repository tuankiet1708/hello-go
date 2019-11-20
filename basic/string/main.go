package string

import (
	"fmt"
)

func Run() {
	var s string = `Hello Go World!\n`
	fmt.Println(s)
	fmt.Println(len(s))
	sl := s[6:8]
	fmt.Println(sl)
	sl = s[6:]
	fmt.Println(sl)
	sl = s[:8]
	fmt.Println(sl)
}
