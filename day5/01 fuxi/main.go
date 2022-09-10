package main

import (
	"encoding/json"
	"fmt"
)

type temp struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func main() {
	var a = struct {
		x int
		y int
	}{x: 2, y: 1}
	fmt.Println(a)

	var a1 temp
	a1.X = 1
	a1.Y = 2

	a2 := temp{
		X: 0,
		Y: 0,
	}
	fmt.Println(a2)

	//调用构造函数
	a3 := newTemp(1, 2)
	fmt.Println(a3)

	a3.dream()
	a3.exchange()
	fmt.Println(a3)

	marshal, err := json.Marshal(a3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))

	s1 := `{"x":2,"y":4}`
	var a4 temp
	err1 := json.Unmarshal([]byte(s1), &a4)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(a4)
}

//构造函数 返回值是对应的结构体类型
func newTemp(x, y int) temp {
	return temp{
		X: x,
		Y: y,
	}
}

//接收者是用对应类型的首字母小写
//指定接收者之后 只有该类型的变量才有资格调用
func (t temp) dream() {
	fmt.Println("temp也有梦想")
	fmt.Println(t.X + t.Y)
}

//指针接收者
//1.需要修改结构体变量的值时需要使用指针接收者
//2.结构体本身比较大，拷贝的内存开销比较大时也要使用指针接收者
//3.保持一致性：如果有一个方法使用了指针接收者，其他的方法为了统一也要使用指针接收者
func (t *temp) exchange() {
	temp := t.X
	t.X = t.Y
	t.Y = temp
}

type addr struct {
	city, province string
}

type student struct {
	name string
	addr //匿名嵌套结构体，就是用类型名字作为名称
}
