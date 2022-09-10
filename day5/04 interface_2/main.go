package main

import "fmt"

//不管是什么牌子的车都能跑

type baoma struct {
}

type benchi struct {
}

func (b baoma) run() {
	fmt.Println("baoma run")
}

func (b benchi) run() {
	fmt.Println("benchi run")
}

//定义一个接口类型 不管是什么结构体 只要是run方法都是car类型
type car interface {
	run()
}

func drive(c car) {
	c.run()
}

func main() {
	var bao baoma
	var ben benchi
	car.run(bao)
	car.run(ben)
	drive(bao)
	drive(ben)
}
