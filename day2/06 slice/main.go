package main

import "fmt"

func main() {
	//切片的定义
	var s1 []int //定义一个存放int类型元素的切片
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//长度和容量
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("len:%d,cap:%d\n", len(s2), cap(s2))

	//2.由数组定义切片
	a := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Println(cap(a))
	b := a[1:4] //[3 5 7] 左闭右开 基于一个数组进行切割
	fmt.Println(b)
	b1 := a[:4] //0-4
	fmt.Println(b1)
	b2 := a[2:] //
	fmt.Println(b2)
	b3 := a[:]
	fmt.Println(b3)

	//切片的长度就是元素的个数，切片的容量就是底层数组从切片第一个元素到最后一个元素的数量
	fmt.Println(len(b), cap(b))
	//3.切片再切片
	b4 := b[1:2] //[5 7] 但是b的容量已经是6了 这时候切的b从第一位切起 那么容量应该是5
	fmt.Println(b4, len(b4), cap(b4))
	fmt.Println(b)
	a[2] = 10
	//这里说明了切片是引用类型，都指向了底层的数组，修改了底层数组，那么上层的切片值肯定会变化
	fmt.Println(b)
	fmt.Println(b4)
}
