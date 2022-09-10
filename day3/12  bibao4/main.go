package main

import "fmt"

func main() {
	f, f2 := calc(10)
	fmt.Println(f(1), f2(2)) //11 9
	fmt.Println(f(3), f2(4)) //12 8
	fmt.Println(f(5), f2(6)) //13 7

}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub

}
