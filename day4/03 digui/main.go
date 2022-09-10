package main

import "fmt"

func main() {
	//递归:自己调用自己
	//递归适合处理那种问题相同但是规模越来越小的场景
	//递归一定要有一个明确的退出条件

	println(Factorial(7))

	fmt.Println(taijie(4))
}

func Factorial(n int) (result int) {
	if n == 1 {
		return 1
	} else {
		result = n * Factorial(n-1)
		return
	}
}

//上台阶面试题
//n个台阶 一次可以走1步 一次可以走2步 有多少种走法
func taijie(n int) (result int) {
	if n == 1 {
		result = 1 //如果只有1个台阶就一种走法
		return
	} else if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}
