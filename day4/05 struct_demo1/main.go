package main

import "fmt"

//结构体

type person struct {
	name   string
	age    int
	hobby  []string
	gender string
}

func main() {
	//声明一个person类型的变量
	var f person
	//通过字段赋值
	f.gender = "男"
	f.hobby = make([]string, 10)
	f.hobby[0] = "football"
	f.hobby[1] = "basketball"
	f.age = 18
	f.name = "ljs"
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	fmt.Println(f.hobby)

	var f1 person
	f1.name = "jwt"
	fmt.Println(f1)

	//匿名结构体  多用于临时场景
	s := struct {
		name string
		age  int
	}{age: 18, name: "fyz"}
	fmt.Println(s)

	var s1 = struct {
		name string
		sex  int
	}{sex: 1, name: "lje"}
	fmt.Println(s1)
}
