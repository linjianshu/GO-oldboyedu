package main

import "fmt"

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("A")
}
func B() {
	//假设此时打开了个数据库连接
	defer func() {
		error := recover()
		fmt.Println(error)
		fmt.Println("要尝试在出错的时候释放数据库连接...")
	}()
	panic("fatal error!") //程序奔溃退出
	fmt.Println("B")
}
func C() {
	fmt.Println("C")
}
