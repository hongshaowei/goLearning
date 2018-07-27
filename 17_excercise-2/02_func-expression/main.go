package main

import "fmt"

func main() {
	half := func(x int) (float64, bool) {
		return float64(x) / 2, x%2 == 0
	}
	fmt.Println(half(5))
	// z, even := half(5)
	// fmt.Println(z, even)
}
