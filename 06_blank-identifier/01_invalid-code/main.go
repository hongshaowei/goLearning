package main

import "fmt"

func main() {
	a := "stored in a"
	fmt.Println(a)
	/*
		whether b is declared but not being used or vice versa are both invalid in golang
		b := "stored in b"
		fmt.Println(b)
	*/
}
