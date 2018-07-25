package main

import (
	"fmt"
)

func main() {
	for i := 60; i < 170; i++ {
		fmt.Printf("%d\t%b\t%#x\t%q\n", i, i, i, i)
	}
}
