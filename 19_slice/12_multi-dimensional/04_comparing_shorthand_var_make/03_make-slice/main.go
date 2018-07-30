package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "James"
	// student = append(student, "James")
	fmt.Println(student)
	fmt.Println(students)
}
