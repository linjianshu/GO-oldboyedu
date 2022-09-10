package main

import (
	"fmt"
)

func main() {
	var n1 int8 = 10
	var n2 int16 = 11
	var n3 int32 = 12
	var n4 int64 = 13
	var n5 = 'a'
	var n6 = '中'
	var f1 float32 = 10.1
	var f2 float64 = 10.2
	var b1 bool = true
	var s1 string = "linjianshu "

	fmt.Printf("类型:%T,值%v\n", n1, n1)
	fmt.Printf("类型:%T,值%v\n", n2, n2)
	fmt.Printf("类型:%T,值%v\n", n3, n3)
	fmt.Printf("类型:%T,值%v\n", n4, n4)
	fmt.Printf("类型:%T,值%c\n", n5, n5)
	fmt.Printf("类型:%T,值%c\n", n6, n6)
	fmt.Printf("类型:%T,值%v\n", f1, f1)
	fmt.Printf("类型:%T,值%v\n", f2, f2)
	fmt.Printf("类型:%T,值%v\n", b1, b1)
	fmt.Printf("类型:%T,值%v\n", s1, s1)

	//统计汉字个数
	s2 := "hello沙河小王子"
	count := 0
	for _, i := range s2 {
		//t:=reflect.TypeOf(i)
		//if reflect.TypeOf(rune('中'))==t {
		//	count++
		//}

		if i > (1 << 7) {
			count++
		}

	}

	fmt.Println(count)

}
