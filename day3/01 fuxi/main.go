package main

import "fmt"

func main() {
	var name string
	name = "ljs"
	fmt.Println(name)
	var ages [30]int
	ages[0] = 1
	ages = [30]int{2, 3, 5}
	fmt.Println(ages)
	ages1 := [...]int{2, 3, 6, 8, 9}
	fmt.Println(ages1)
	ages2 := [...]int{1: 1, 99: 99}
	fmt.Println(ages2)

	//二维数组
	a := [3][2]string{}
	a[0][1] = "ljs"
	a[0][0] = "jwt"
	fmt.Println(a)

	//多维数组是值类型
	a1 := [3][2]string{{"ljs", "jwt"}, {"fyz", "lje"}}
	fmt.Println(a1)

	var a2 = [3][2]int{[2]int{1, 2}, [2]int{3, 4}}
	fmt.Println(a2)

	//数组是值类型
	a3 := [3]int{1, 2, 3}
	fmt.Println(a3)
	f1(a3)
	fmt.Println(a3)

	a4 := []int{1, 2, 3}
	fmt.Println(a4)
	f2(a4)
	fmt.Println(a4)

	//切片
	a5 := []int{}
	fmt.Println(a5)
	fmt.Println(a5 == nil)
	//没有分配内存 零切片声明 nil
	var a6 []int
	fmt.Println(a6)
	fmt.Println(a6 == nil)
	//make初始化 分配内存
	a7 := make([]int, 5, 5)
	fmt.Println(a7)
	fmt.Println(a7 == nil)

	s1 := []int{1, 2, 3}
	s2 := s1
	fmt.Println(s1)
	s2[1] = 100
	fmt.Println(s2)
	fmt.Println(s1) //切片不存值 指向同一个数组

	var s3 []int
	//append将自动初始化分配内存+扩容
	s3 = append(s3, 1)
	fmt.Println(s3)

	var s4 []int
	s4 = make([]int, 1, 1)
	copy(s4, s3) //copy函数必须先将dest切片声明好并且初始化好分配好内存和长度
	fmt.Println(s4)

	//指针
	//go里面的指针只能读不能修改
	addr := "沙河"
	addrpointer := &addr
	fmt.Println(addrpointer)
	fmt.Printf("%T\n", addrpointer)
	fmt.Printf(*addrpointer)
	fmt.Println()

	//map
	var m map[string]int
	m = make(map[string]int, 5)
	m["ljs"] = 99
	m["jwt"] = 98
	fmt.Println(m)
	fmt.Println(m["jiwuming"]) //如果不存在key ，返回的将是value类型的默认值
	score, ok := m["jiwuming"]
	if ok {
		println(score)
	} else {
		println("查无此人")
	}

	delete(m, "lalala") //如果没有的话，什么都不干，不报错
	delete(m, "jwt")
	fmt.Println(m)
}
func f1(a [3]int) {
	//go语言中函数传递的都是值 ctrl+c ctrl+v  感觉这话说的有点问题
	a[1] = 100
}

func f2(a []int) {
	a[1] = 100
}
