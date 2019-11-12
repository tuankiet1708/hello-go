package main

import (
	"fmt"

	"./ifswitchloop"
)

func main() {
	// intstring.Run()

	// arraystruct.RunArray()
	// arraystruct.RunStruct()

	// var studentA arraystruct.Student
	// fmt.Println(studentA)
	// arraystruct.RunStruct()

	// pointer.RunPointer()

	// slicemap.RunSlice()
	// slicemap.RunSliceAdvanced()
	// slicemap.RunMap()

	// function.RunFunc()

	ifswitchloop.RunIfSwitchLoop()
	var fact uint = 1
	fact = ifswitchloop.RunFact1(4)
	fmt.Println(fact)
	fact2 := ifswitchloop.RunFact2(4)
	fmt.Println(fact2)
	fact2 = ifswitchloop.RunFact3(4)
	fmt.Println(fact2)
	fact2 = ifswitchloop.RunFact4(4)
	fmt.Println(fact2)
}
