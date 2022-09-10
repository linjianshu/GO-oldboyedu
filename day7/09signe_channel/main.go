package main

import "fmt"

func main() {
	ch1 = make(chan int, 100)
	ch2 = make(chan int, 100)
	go counter(ch1)
	go squarer(ch1, ch2)
	printer(ch2)
}

var ch1 chan int
var ch2 chan int

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	defer close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}
	defer close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
