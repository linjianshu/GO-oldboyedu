package main

import "fmt"

func main() {
	lixiang(yuanshuai, "ljs")
	f := zhoulin()
	println(f(1, 2))

	low(transfer("ljs"))
	low(transfer1(yuanshuai, "yuanshuai"))
}

func zhoulin() func(int, int) int {
	return func(i int, i2 int) int {
		return i + i2
	}
}

//闭包应用 把yuanshuai放到low里执行
func transfer(name string) func() {
	return func() {
		yuanshuai(name)
	}
}

//闭包应用
func transfer1(f func(a string), name string) func() {
	return func() {
		f(name)
	}
}

func yuanshuai(name string) {
	fmt.Println("hello", name)
}

func low(f func()) {
	f()
}
func lixiang(f func(string2 string), name string) {
	f(name)
}
