package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
func add() {
	//lock.Lock()  //加锁也可以实现
	//x++
	//lock.Unlock()

	atomic.AddInt64(&x, 1) //
	wg.Done()
}
