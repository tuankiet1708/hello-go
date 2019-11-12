package ifswitchloop

import (
	"fmt"
	"time"
)

func RunIfSwitchLoop() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Hôm nay là cuối tuần!")
	default:
		fmt.Println("Hôm nay là ngày làm việc.")
	}
}

func RunFact1(n uint) uint {
	var factn uint = 1
	var i uint = 1

	for {
		factn *= i
		i++
		if i > n {
			break
		}
	}

	return factn
}

func RunFact2(n uint) uint {
	var i, factn uint = 1, 1

	for i <= n {
		factn *= i
		i++
	}

	return factn
}

func RunFact3(n uint) uint {
	var factn uint = 1

	for i := uint(1); i <= n; i++ {
		factn *= i
	}
	return factn
}

func RunFact4(n uint) uint {
	var factn uint = 1
	a := make([]int, n)

	for i, _ := range a {
		factn *= uint(i + 1)
	}

	return factn
}
