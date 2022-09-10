package main

import "fmt"

var name string
var age int
var isOk bool

//批量声明
var (
	name1 string
	age1  int
	isOk1 bool
)

func main() {
	fmt.Println(name1)
	fmt.Println(age1)
	fmt.Println(isOk1)

	name1 = "linjianshu"
	age1 = 16
	isOk1 = true

	//go语言中推荐使用驼峰命名
	//go语言中非全局变量声明必须使用，不用就编译不过去
	fmt.Println(name1)
	fmt.Println(age1)
	fmt.Println(isOk1)
	fmt.Printf("name:%s", name1) //%s占位符，使用name1这个变量去替换这个占位符

	//声明变量同时赋值
	var studentName string = "ljs"
	//类型推导
	var studentName1 = "ljs"
	fmt.Println(studentName)
	fmt.Println(studentName1)
	//简短变量声明 :=  只能在函数里面使用
	studentName2 := "jwt"
	fmt.Println(studentName2)

	x, _ := foo()
	fmt.Println(x)

	x1 := 0
	x1, _ = foo()
	fmt.Println(x1)

}

func foo() (int, string) {
	return 10, "ljs"
}
