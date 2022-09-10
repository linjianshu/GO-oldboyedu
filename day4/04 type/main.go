package main

import "fmt"

//type后面跟的是类型
type myInt int     //自定义类型
type yourInt = int //类型别名

func main() {
	//自定义类型和类型别名
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var r rune
	r = '中'
	fmt.Printf("%c\n", r)
	fmt.Printf("%T\n", r)
}
