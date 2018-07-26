package main

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	sum := 0
	for i := 1; i < b.N; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

// run this at command:
// go test -bench='.*'
