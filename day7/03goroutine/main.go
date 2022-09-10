package main

import (
	"fmt"
	"time"
)

//程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 10; i++ {
		//go hello(i)  //开启一个单独的goroutine去执行hello函数（任务）

		//匿名
		//go func() {
		//	fmt.Println(i)
		//}()  //闭包 会出现i外层是1 内部输出1 外部这时候已经跑到10了 那么这时候内部就输出10

		//
		go func(i int) {
			fmt.Println(i) //用的是函数参数的那个i ， 不再是外面的那个i了
		}(i)

	}
	fmt.Println("main")
	//main函数结束了 由main函数启动的goroutine也都结束了
	time.Sleep(time.Second)
}

func hello(i int) {
	fmt.Println("hello", i)
}
