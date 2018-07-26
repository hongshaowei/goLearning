package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	// var b *int = &a, still syntax checker suggest to omit the type and let go compiler to infer it
	var b = &a
	fmt.Println(b)
	// the above code made b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -> '*' is part of the type -- b is of type *int
}
