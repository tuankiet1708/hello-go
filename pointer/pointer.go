package pointer

import "fmt"

func RunPointer() {
	i := 10
	fmt.Println(i)  //  10
	fmt.Println(&i) //   0xc082008340
	var p1 *int
	fmt.Println(p1) //   nil
	p1 = &i
	p2 := &i
	fmt.Println(p1)  //  0xc082008340
	fmt.Println(p2)  //  0xc082008340
	fmt.Println(&p1) //  0xc082024020

	fmt.Println(*p2) // 10
	*p2 = 5
	fmt.Println(i)   // 5
	fmt.Println(*p2) // 5
	fmt.Println(*p1) // 5

	var p3 = new(int)
	fmt.Println(p3)  //  0xc082008340
	fmt.Println(*p3) //  0
	*p3 = 10
	fmt.Println(*p3) //  10
}
