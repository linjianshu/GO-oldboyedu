package main

import "fmt"

type myInt int

//结构体遇到的问题
//1.myInt(100)是什么

//问题2： 结构体初始化
type person struct {
	name string
	age  int
}

//问题3：为什么要有构造函数
func newPerson(name string, age int) person {
	return person{name: name, age: age}
}

func main() {
	//声明int32的值
	var a1 int32
	a1 = 10
	fmt.Println(a1)

	var a2 int32 = 10
	fmt.Println(a2)

	a3 := int32(10)
	fmt.Println(a3)

	var a4 = int32(10)
	fmt.Println(a4)

	//自定义类型同上
	var b1 myInt
	b1 = 100
	fmt.Println(b1)

	var b2 myInt = 100
	fmt.Println(b2)

	b3 := myInt(100)
	fmt.Println(b3)

	var b4 = myInt(100)
	fmt.Println(b4)

	var p1 person
	p1.name = "ljs"
	p1.age = 18

	p2 := person{
		name: "ljs",
		age:  19,
	}
	fmt.Println(p2)

	p3 := person{
		"jwt",
		10,
	}
	fmt.Println(p3)

	s1 := []string{"100", "200"}
	fmt.Println(s1)
	m1 := map[int]string{
		1: "100",
		2: "200",
	}
	fmt.Println(m1)
}
