package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `{"name":"ljs","age":16}`
	var p person
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		return
	}
	fmt.Println(p)
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
