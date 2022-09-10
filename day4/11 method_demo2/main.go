package main

import "fmt"

//给自定义类型添加方法
//不能给别的包里面的类型添加方法，只能给自己的包里的类型添加方法
type myInt int

func (i myInt) hello() {
	fmt.Println("this is a int" + (string(i)))
}
func main() {
	var i myInt
	i = 10
	i.hello()

}
