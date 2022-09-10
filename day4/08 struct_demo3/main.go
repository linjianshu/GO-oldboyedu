package main

import "fmt"

type x struct {
	a, b, c int8
}

func main() {
	//结构体占用一块连续的内存空间
	x := x{
		a: 10,
		b: 20,
		c: 30,
	}

	fmt.Printf("%p\n", &(x.a))
	fmt.Printf("%p\n", &(x.b))
	fmt.Printf("%p\n", &(x.c))
}
