package main

import (
	"fmt"
)

func main() {
	//map和slice的组合
	a := []map[string]int{}
	//元素类型为map的切片
	var a1 = make([]map[string]int, 10, 10)
	//没有对内部的map做初始化
	a1[0] = make(map[string]int, 10)
	a1[0]["ljs"] = 9
	a1[0]["jwt"] = 8
	fmt.Println(a)
	fmt.Println(a1)

	//值为切片类型的map
	var a2 = make(map[string][]string, 10)
	a2["ljs"] = make([]string, 10, 10)
	a2["ljs"] = []string{"giegie"}
	fmt.Println(a2)
	a2["ljs"] = append(a2["ljs"], []string{"jiejie", "didi"}...)
	fmt.Println(a2)

	s1 := "how do you do"

	idx := []int{0}
	for i := 0; i < len(s1); i++ {
		if s1[i] == ' ' {
			idx = append(idx, i)
		}
	}
	fmt.Println(idx)
	for i := 0; i < len(idx)-1; i++ {
		fmt.Println(s1[idx[i]:idx[i+1]])
	}
}
