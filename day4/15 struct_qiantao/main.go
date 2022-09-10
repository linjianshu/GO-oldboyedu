package main

import "fmt"

type person struct {
	name string
	age  int
	addr address
}

type company struct {
	name    string
	address //匿名嵌套结构体 可以直接拿到匿名结构体里面的字段
}

type address struct {
	province string
	city     string
}

func main() {
	p1 := person{
		name: "ljs",
		age:  18,
		addr: address{city: "fuzhou", province: "fujian"},
	}
	fmt.Println(p1.addr.province)

	c1 := company{
		name:    "alibaba",
		address: address{province: "zhejiang", city: "hangzhou"},
	}
	fmt.Println(c1.city) //先在自己结构体找这个字段 找不到就去匿名嵌套的结构体中查找该字段
}
