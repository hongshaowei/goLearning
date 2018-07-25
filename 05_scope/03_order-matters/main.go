package main

import "fmt"

func main() {
	// x is not recognized before it is initialized
	// fmt.Println(x)
	fmt.Println(y)
	x := 42
	fmt.Println(x)
}

var y = 42
