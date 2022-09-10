package main

import (
	"fmt"
	"sync"
	"time"
)

var i = make(chan bool, 1)
var wg sync.WaitGroup

//为什么需要context
func main() {
	wg.Add(1)
	go f()
	//如何通知子goroutine退出
	time.Sleep(time.Second)
	i <- true
	wg.Wait()
}

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("ljs")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-i:
			break LOOP
		default:

		}
	}
}
