package main

import "fmt"

func main() {

	b := true

	// the initialization part; the if statement expression {}
	// initialization part is just to initialize a var, the expression part is the place to decide
	// and the var is in the if block
	if food := "Chocolate"; b {
		fmt.Println(food)
	}

}
