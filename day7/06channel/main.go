package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int //需要指定通道中元素的类型
var wg sync.WaitGroup

func main() {
	//noBufChannel()
	withBufChannel()
}

func f1() {
	fmt.Println("后台goroutine从通道b中取到了：", <-b)
	defer wg.Done()
}
func withBufChannel() {
	fmt.Println(b)
	b = make(chan int, 16) //带缓冲区通道的初始化
	b <- 10
	fmt.Println(b)
	fmt.Println(<-b)
	close(b)
}
func noBufChannel() {
	wg.Add(1)
	fmt.Println(b) //nil
	go f1()
	b = make(chan int) //不带通道的初始化
	b <- 10
	fmt.Println(b)
	wg.Wait()
}
