package main

import (
	"fmt"
)

type person struct {
	name string
	sex  string
}

func main() {
	//结构体是值类型
	p := person{
		name: "ljs",
		sex:  "男",
	}
	fmt.Println(p)

	var p1 person
	p1.name = "lje"
	p1.sex = "nan"

	var p2 person
	p2 = p1
	p2.name = "fyz"
	fmt.Println(p2)
	fmt.Println(p1)

	func(x person) {
		x.sex = "女" //传的是值
	}(p2)
	fmt.Println(p2)

	func(x *person) {
		(*x).sex = "nv" //传的是地址
		//x.sex = "nv"  //语法糖 一样的同上
	}(&p2)
	fmt.Println(p2)

	//创建一个指针类型的person
	var p3 = new(person) //new 返回的是指针地址 这个类型
	p3.sex = "nan"
	(*p3).name = "www" //语法糖 一样的同上
	fmt.Println(p3)
	fmt.Printf("%T\n", p3)
	fmt.Printf("%p\n", p3) //返回的是这个指针的值 p3保存的值就是一个内存地址
	fmt.Printf("%v\n", p3)
	fmt.Printf("%T\n", &p3)
	fmt.Printf("%p\n", &p3) //返回的是这个指针类型的值的地址

	//key value 初始化
	var p4 = &person{
		name: "lll",
	}
	fmt.Println(p4)

	//使用值 列表的形式初始化 顺序保持一致
	p5 := person{
		"nv",
		"slkdjf",
	}
	fmt.Println(p5)
}
