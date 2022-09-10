package main

import (
	"fmt"
	"time"
)

var jobs chan int
var results chan int

func main() {
	jobs = make(chan int, 100)
	results = make(chan int, 100)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	for i := 0; i < 5; i++ {
		<-results
	}
	//for result := range results {
	//	fmt.Println(result)
	//}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " start job ", j)
		time.Sleep(time.Second)
		results <- 2 * j
		fmt.Println("worker ", id, " end job", j)
	}
}
