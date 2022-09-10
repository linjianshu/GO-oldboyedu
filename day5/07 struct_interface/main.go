package main

import "fmt"

//接口还可以嵌套
type animal interface {
	mover
	eater
}

//同一个结构体可以实现多个接口
type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet byte
}

//一个结构体可以实现多个接口
func (c *cat) move() {
	fmt.Println(c.name + " is moving")
}

func (c *cat) eat(something string) {
	fmt.Println(c.name + " is eating " + something)
}
func main() {
	c1 := cat{
		name: "tom",
		feet: 4,
	}
	mover.move(&c1)
	eater.eat(&c1, "猫粮")
}
