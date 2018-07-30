package main

import "fmt"

func main() {
	n := hashBucket("Go", 12)
	fmt.Println(n)
}

func hashBucket(word string, buckets int32) int32 {
	// letter := int(word[0])
	// Here to get the first letter of the string, so the type should be rune
	letter := rune(word[0])
	bucket := letter % buckets
	return bucket
}
