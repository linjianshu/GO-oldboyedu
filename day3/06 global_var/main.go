package main

import "fmt"

//变量的作用域
var a = 100

func main() {
	f1()
}

func f1() {
	fmt.Println(a)
}
