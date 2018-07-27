package main

import (
	"fmt"
)

func half(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}

func main() {
	func() {
		z, even := half(5)
		fmt.Println(z, even)
	}() // writting a self-executed function remember the parens at the end of it
}
