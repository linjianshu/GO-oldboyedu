package main

import (
	"fmt"
)

func main() {
	//函数
	i := sum(10, 20)
	fmt.Println(i)

	f1(10, 20)

	f2()

	fmt.Println(f3())

	i2, i3 := f4()
	fmt.Println(i2, i3)

	f6(5, "ljs", "jwt", "lje", "fyz")
}

//有参数有返回值 返回值定义变量
func sum(x int, y int) (ret int) {
	ret = x + y
	return ret
}

//有参数无返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//无参数无返回值
func f2() {
	fmt.Println("hello world!")
}

//无参数有返回值
func f3() string {
	return "hello world"
}

//无参数有多个返回值
func f4() (int, int) {
	return 1, 2
}

//参数类型简写
func f5(x, y int, m, n string) {
	fmt.Println(x + y)
}

//可变长度函数
func f6(x int, y ...string) {
	fmt.Println(x, y) //y是切片 切片类型由形参定义
}

//go语言中没有默认参数这个概念
