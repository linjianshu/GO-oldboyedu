package main

import "fmt"

//匿名字段
type person struct {
	string
	int
}

func main() {
	a := person{
		"ljs",
		10,
	}
	fmt.Println(a.string)
	fmt.Println(a.int)
}
