package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Please enter two number:")
	fmt.Scan(&a, &b)
	if !(a == 0 || b == 0) {
		if a > b {
			fmt.Println(a % b)
		} else {
			fmt.Println(b % a)
		}
	} else {
		fmt.Println("Please enter numbers not equal to zero")
	}
}
