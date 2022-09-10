package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//waitGroup
	//f1()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f2(i)
	}

	//如何知道这10个goroutine都结束了
	wg.Wait() //等待wg的计数器减为0
}

var wg sync.WaitGroup

func f1() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		i1 := rand.Int()
		i2 := rand.Intn(10) //左开右闭
		fmt.Println(i1, i2)
	}
}

func f2(i int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
	defer wg.Done()
}
