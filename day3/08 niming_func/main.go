package main

import "fmt"

func main() {
	println(a(10))
	//但是通常匿名函数不是这么用的 通常是由于函数内部不允许定义函数，所以使用匿名函数现写现用
	a := func(x, y int) int {
		return x + y
	}
	println(a(10, 20))
	//如果只是调用一次的函数，还可以简写成立即执行函数
	i := func(x, y int) int {
		return x * y
	}(10, 20)
	fmt.Println(i)
}

var a = func(x int) int {
	return x
}
