package main

import (
	"fmt"
)

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = []int{} //这样声明没办法复制进去
	var a4 []int     //这样声明也没办法复制进去
	var a5 = make([]int, len(a1), cap(a1))
	copy(a3, a1)
	copy(a4, a1)
	copy(a5, a1)
	fmt.Println(a1, a2, a3, a4, a5)

	a1[0] = 100
	fmt.Println(a1, a2, a3, a4, a5)

	//删除第二个元素
	a5 = append(a5[:1], a5[2:]...)
	fmt.Println(a5)
	fmt.Println(cap(a5))

	//验证
	//1.切片不保存具体的值
	//2.切片对应一个底层数组
	//3.底层数组都是占用一块连续的内存
	x1 := [...]int{1, 3, 5} //数组
	x2 := x1[:]             //切片 切片指向底层数组
	fmt.Println(x2, len(x2), cap(x2))
	fmt.Printf("%p\n", &x1[0])
	x2 = append(x1[:1], x1[2:]...) //切片截取底层数组 重新定义了底层数组的索引的值
	fmt.Printf("%p\n", &x2[0])     //说明指向的底层数组地址没变 变了的是地址里的值
	fmt.Println(x2)                //切片的索引的值
	fmt.Println(x1)                //被修改后的底层数组的索引和值
}
