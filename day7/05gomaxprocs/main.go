package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2) //默认CPU的逻辑核心数，默认跑满整个CPU
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
	defer wg.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
	defer wg.Done()
}
