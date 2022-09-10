package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%T \n", f1)
	fmt.Printf("%T \n", f2)
	fmt.Printf("%T \n", f3)
	f3(f2)

	a := f4(f2)
	fmt.Printf("%T\n", a)
}

func f1() {
	fmt.Println("hello 沙河")
}

func f2() int {
	return 5
}

//函数也可以作为参数的类型
func f3(x func() int) {
	fmt.Println(x())
}

//函数还可以作为返回值的类型
func f4(x func() int) func(int, int) int {
	return f5
}

func f5(x, y int) int {
	return x + y
}
