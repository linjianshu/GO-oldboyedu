package main

import "fmt"

//go语言中的函数的return不是原子操作，在底层是分为两步来执行
//第一步：返回值赋值
//第二步：真正的return返回
//函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
func main() {
	fmt.Println(f1())   //5
	fmt.Println(f2())   //6
	fmt.Println(f3())   //5
	fmt.Println(f3_1()) //[100 2]
	fmt.Println(f4())
	fmt.Println(f5())
}

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值是x x又++了 所以返回6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ //是因为int是值类型 所以y是拷贝值而不是拷贝地址的原因吗
	}()
	return x
}

func f3_1() (y []int) {
	x := []int{1, 2}
	defer func() {
		x[0] = 100 //因为[]int 切片是引用类型 所以y拷贝的是地址而不是值
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++ //改变的是函数的副本
	}(x) //(x)代表的是传入参数
	return 5
}

func f5() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5
}
