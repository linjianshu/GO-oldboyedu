package main

import "fmt"

func main() {
	//数组
	//存放元素的容器
	//必须指定存放的元素的类型和容量（长度）
	//数组的长度是数组类型的一部分 也就是尽管类型一致但是长度不一致也不是同一个数组类型
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:type%T, a2:type%T", a1, a2)
	fmt.Println()

	//数组的初始化
	//如果不初始化：默认元素都是零值(布尔值就是false，整型和浮点型都是0，字符串就是“”)
	fmt.Println(a1, a2)
	//1.初始化方式1
	b1 := [3]bool{true, true, true}
	fmt.Println(b1)
	//2.初始化方式2 根据初始值自动推断数组的长度是多少
	b2 := [...]int{1, 3, 4, 2, 6, 2, 73, 12}
	fmt.Println(b2)
	fmt.Println(len(b2))
	//3.初始化方式3 根据索引初始化
	b3 := [5]int{1, 2}
	fmt.Println(b3)
	b3 = [5]int{0: 1, 4: 2}
	fmt.Println(b3)

	//数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	//1.for range
	for _, v := range citys {
		fmt.Println(v)
	}

	for i, _ := range citys {
		fmt.Println(citys[i])
	}

	//2.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	//多维数组
	c1 := [3][2]int{0: [2]int{2, 3}, 1: [2]int{4, 5}}
	fmt.Println(c1)
	c2 := [3][2]int{{1, 2}, {3, 4}}
	fmt.Println(c2)

	//多维数组的遍历
	//var b11:=[2][3]string{{"a","b","c"},{"d","e","f"}}
	//可以这么记 go语言中 实际数组展示使用空格来区分的，但是声明的时候需要用逗号隔开
	for _, v := range c1 {
		fmt.Println(v)
		for _, v1 := range v {
			fmt.Printf("%d ", v1)
		}
		fmt.Println()
	}

	for i := 0; i < len(c2); i++ {
		fmt.Println(c2[i])
		for j := 0; j < len(c2[j]); j++ {
			fmt.Printf("%d ", c2[i][j])
		}
		fmt.Println()
	}

	//数组是值类型
	d1 := [...]int{1, 2, 3}
	d2 := d1
	d2[0] = 100
	fmt.Println(d1)
	fmt.Println(d2)

	//练习
	e := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range e {
		sum += v
	}
	fmt.Println(sum)

	for i, _ := range e {
		for j := i + 1; j < len(e); j++ {
			if e[i]+e[j] == 8 {
				fmt.Printf("(%d %d)", i, j)
				break
			}
		}
		fmt.Println()
	}
}
