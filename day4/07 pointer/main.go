package main

import "fmt"

func main() {
	var a int
	a = 100
	b := &a
	fmt.Printf("%T %p\n", &a, &a)
	fmt.Printf("%T %p\n", b, b)   //b的值
	fmt.Printf("%T %v\n", b, b)   //b的值
	fmt.Printf("%T %p\n", &b, &b) //b的内存地址
}
