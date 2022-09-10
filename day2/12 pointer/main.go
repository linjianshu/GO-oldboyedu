package main

import "fmt"

func main() {
	//1.&取地址
	//2.*根据地址取值
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p) //*int表示int类型的指针
	m := *p
	fmt.Printf("%v\n", m)
	fmt.Printf("%T\n", m)

	var a *int
	fmt.Println(a)    //nil 赋值会报错 空指针异常
	var a1 = new(int) //使用new关键字会分配内存块 不会造成空指针
	fmt.Println(a1)
	fmt.Println(*a1)
	*a1 = 100
	fmt.Println(*a1)

}
