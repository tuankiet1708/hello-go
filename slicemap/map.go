package slicemap

import "fmt"

func RunMap() {
	var student map[int]string
	fmt.Println(student) // map[]
	// student[1] = "Mai"   // Go báo lỗi do chưa cấp phát cho student
	student = make(map[int]string)
	student[1] = "Mai"
	student[2] = "Lan"
	student[3] = "Cúc"
	student[4] = "Trúc"
	age := map[string]int{
		"Mai":  20,
		"Lan":  18,
		"Cúc":  21,
		"Trúc": 16,
	}
	fmt.Println(student) // map[4:Trúc 1:Mai 2:Lan 3:Cúc]
	fmt.Println(age)     // map[Lan:18 Cúc:21 Trúc:16 Mai:20]
	i := 1
	fmt.Println("SV 1: ", student[i], " - ", age[student[i]]) // SV 1: Mai - 20
	i++
	delete(age, student[i])
	fmt.Println("SV 2: ", student[i], " - ", age[student[i]]) // SV 2: Lan - 0
	age["Cúc"]++
	fmt.Println("SV 3: ", student[3], " - ", age[student[3]]) // SV 3: Cúc - 22
	age[student[4]] = 23
	fmt.Println("SV 4: ", student[4], " - ", age[student[4]]) // SV 4: Trúc - 23
	age["trúc"] = 22
	fmt.Println("SV 4: ", student[4], " - ", age[student[4]]) // SV 4: Trúc - 23
	fmt.Println(age)                                          // map[Mai:20 trúc:22 Cúc:22 Trúc:23]
}
