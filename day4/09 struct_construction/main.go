package main

import "fmt"

type person struct {
	name string
	age  int
}

//构造函数 以new开头 约定成俗
//构造函数返回结构体还是返回结构体指针是有考量的
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson1(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("ljs", 18)
	p2 := newPerson("jwt", 20)
	fmt.Println(p1)
	fmt.Println(p2)

	p3 := newPerson1("ljs", 18)
	p4 := newPerson1("jwt", 20)
	fmt.Printf("%T %v \n", p3, p3)
	fmt.Printf("%v\n", *p3)
	fmt.Printf("%T %v \n", p4, p4)
	fmt.Printf("%v\n", *p4)

}
