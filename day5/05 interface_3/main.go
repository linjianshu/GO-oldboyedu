package main

import "fmt"

//接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int
}

type hen struct {
	feet int
}

func (h hen) move() {
	fmt.Println("鸡动!")
}

func (h hen) eat(a string) {
	fmt.Println("吃鸡饲料..." + a)
}

func (c cat) move() {
	fmt.Println(c.name)
}

func (c cat) eat(a string) {
	fmt.Println("猫吃" + a)
}

func eat(a animal, s string) {
	a.eat(s)
}
func main() {

	var a animal //定义一个接口类型的变量
	c := cat{
		name: "ljs",
		feet: 4,
	}
	a = c
	a.eat("mao")
	fmt.Printf("%T\n", a)

	h := hen{feet: 2}
	a = h
	a.eat("和你")
	fmt.Printf("%T\n", a)

	eat(c, c.name)
	c1 := animal(c)
	c1.eat(c.name)
}
