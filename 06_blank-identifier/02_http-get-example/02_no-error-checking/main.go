package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// http.Get returns respond and error,
	// but I don't want to handle error or do anything with it further
	// so just use _ represents a blank idetifier
	res, _ := http.Get("http://www.geekwiseacademy.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
