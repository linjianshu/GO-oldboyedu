package main

import "fmt"

func main() {
	//闭包是什么
	//闭包是一个函数，这个函数包含了他外部作用域的一个变量
	//底层
	//1.函数可以作为返回值
	//2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找
	ret := adder(100)
	i := ret(200)
	fmt.Println(i)
}

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
