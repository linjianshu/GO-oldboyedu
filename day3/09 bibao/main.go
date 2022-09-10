package main

import "fmt"

func main() {
	f1(f3(1, 2))
	f1(f4(1, 2))
	f1(f5(f2, 1, 2))
}

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//如何让f1调用的时候执行f2 也就是两个同事写的代码相互兼容
//由于f1的形参是一个无形参无返回值的函数类型，因此需要构造一个函数，让其返回值是无形参无返回值的函数类型f3
//当然为了兼容f2，f3的形参需要和f2的形参相匹配，这样一来在执行f3的时候，内部调用了f2，并且返回类型满足f1所需
func f3(x, y int) func() {
	func(x, y int) {
		f2(x, y)
	}(x, y)
	return func() {
	}
}

func f4(x, y int) func() {
	return func() {
		f2(x, y)
	}
}

func f5(f func(int, int), x, y int) func() {
	//把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	return func() {
		f(x, y)
	}
}
