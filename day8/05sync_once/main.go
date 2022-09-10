package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once
var ch1 = make(chan int, 100)
var ch2 = make(chan int, 100)

func main() {
	wg.Add(3)
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()
	for i := range ch2 {
		fmt.Println(i)
	}
}

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() {
		close(ch2)
	})
}
