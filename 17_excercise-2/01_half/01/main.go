package main

import (
	"fmt"
)

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	z, even := half(4)
	fmt.Println(z, even)
}
