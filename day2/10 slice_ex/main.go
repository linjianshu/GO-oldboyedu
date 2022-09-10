package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	b := [...]int{3, 7, 8, 9, 1}
	b1 := b[:]
	sort.Ints(b1) //对切片进行排序
	fmt.Println(b1)
}
