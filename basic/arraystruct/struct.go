package arraystruct

import "fmt"

type Student struct {
	id   int
	Name string
	Age  int
}

func RunStruct() {
	var studentA Student
	fmt.Println(studentA) // {0 "" 0}
	studentA.id = 1
	studentA.Name = "Mai"
	studentA.Age = 20
	var studentB = Student{id: 2, Name: "Lan", Age: 18}
	var studentC = Student{id: 3, Name: "Cúc"}
	studentD := Student{4, "Trúc", 16}
	fmt.Println(studentA) // {1 Mai 20}
	fmt.Println(studentB) // {2 Lan 18}
	fmt.Println(studentC) // {3 Cúc 0}
	fmt.Println(studentD) // {4 Trúc 16}
}
