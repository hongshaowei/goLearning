package main

import "fmt"

// var declared outside func is all-visible in this package,
// ps: if var declared outside func needs keyword var
var x = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
