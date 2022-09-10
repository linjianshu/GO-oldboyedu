package main

import "fmt"

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Println(string(a.name) + "会动")
}

type dog struct {
	feet byte
	animal
}

func (d dog) wang() {
	fmt.Println(d.name + "wangwangwang")
}

func newDog(a animal, feet byte) dog {
	return dog{
		feet, a,
	}
}

func newAnimal(name string) animal {
	return animal{name: name}
}
func main() {
	//结构体模拟实现其他语言中的继承
	newDog(newAnimal("jwt"), 4).wang()

	d1 := dog{
		4, animal{name: "ljs"},
	}
	d1.move() //只能匿名嵌套结构体才能实现类似于继承的效果 如果有名字好像就调用不了

}
