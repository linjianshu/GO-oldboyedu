package main

import "fmt"

var ch1 chan int
var ch2 chan int

func main() {
	ch1 = make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	ch2 = make(chan int)
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}
