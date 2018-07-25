package main

import "fmt"

func main() {
	var message string
	message = "Let's Go!"

	var a, b, c int
	a = 1

	fmt.Printf("%v\t%T\n", message)
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
}
