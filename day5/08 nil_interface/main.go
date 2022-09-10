package main

import "fmt"

//空接口
func main() {
	m1 := make(map[interface{}]interface{}, 10)
	m1[1] = "hello world"
	m1["hello world"] = 1
	m1[false] = [...]string{"1", "2", "3"}
	m1[[...]int{1, 2}] = []bool{true, false}
	fmt.Println(m1)

	show(m1)
}

func show(a interface{}) {
	fmt.Printf("%T  %v\n", a, a)
}
