package main

import "fmt"

//引出接口的实例

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("miaomiaomiao~")
}

func (d dog) speak() {
	fmt.Println("wangwangwang~")
}

type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}

func fuck(a speaker) {
	a.speak()
}

func main() {
	c := cat{}
	d := dog{}
	fuck(c)
	fuck(d)

	var ss1 speaker //定义一个接口类型 ：speaker的变量
	ss1 = d
	ss1.speak()

	ss := speaker(c)
	ss.speak()
}
