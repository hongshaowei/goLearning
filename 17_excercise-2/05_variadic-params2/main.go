package main

import "fmt"

func main() {
	foo := func(x ...int) {
		fmt.Println(x)
	}
	foo()
	foo(1, 2)
	iSlice := []int{1, 2, 3, 4}
	foo(iSlice...)
}
