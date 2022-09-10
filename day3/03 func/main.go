package main

import "fmt"

func main() {
	println(f4(1, 2, 3, 4))
}

func f1() {
	fmt.Println("hello 沙河")
}

func f2(name string) {
	fmt.Println("hello", name)
}

func f3(x, y int) int {
	return x + y //y是一个可变长度的切片类型
}

func f4(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func f5(x, y int) (sum int) {
	sum = x + y
	return
}

func f6(x, y int) (x1, y1 int) {
	x1 = x
	y1 = y
	return
}
