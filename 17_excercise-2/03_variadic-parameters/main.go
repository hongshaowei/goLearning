package main

import (
	"fmt"
)

func main() {
	max := func(x ...int) int {
		var z int
		for _, v := range x {
			if v > z {
				z = v
			}
		}
		return z
	}
	fmt.Println(max(1, 2, 3, 4, 5))
}
