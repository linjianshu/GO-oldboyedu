package main

import (
	"fmt"
)

//使用值接收者和指针接收者的区别
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int
}

////使用值接收者实现了接口的所有方法
//func (c cat) move() {
//	fmt.Println("走猫步")
//}
//
//func (c cat) eat(a string) {
//	fmt.Println("猫吃"+a)
//}

//使用指针接收者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("走猫步")
}

func (c *cat) eat(a string) {
	fmt.Println("猫吃" + a)
}

func main() {
	var a1 animal
	c1 := cat{
		name: "tom",
		feet: 4,
	}
	c2 := &cat{
		name: "假老练",
		feet: 4,
	}

	a1 = &c1
	fmt.Println(a1)
	a1.eat("bianbian")

	a1 = c2
	fmt.Println(a1)
	a1.eat("大便便")

}
