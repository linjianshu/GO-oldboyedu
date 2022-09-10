package main

import (
	"fmt"
	"sync"
	"time"
)

//第一种就是通过全局变量 来控制goroutine退出
var notify bool
var wg sync.WaitGroup

//为什么需要context
func main() {
	wg.Add(1)
	go f()
	//如何通知子goroutine退出
	time.Sleep(time.Second)
	notify = true
	wg.Wait()
}

func f() {
	defer wg.Done()
	for !notify {
		fmt.Println("ljs")
		time.Sleep(time.Millisecond * 500)
	}

}
