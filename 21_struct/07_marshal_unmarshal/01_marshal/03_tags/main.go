package main

import (
	"encoding/json"
	"fmt"
)

// tags will replace the original field name as the key in json file
// "-" means skip this field when formatting to json format
type person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
