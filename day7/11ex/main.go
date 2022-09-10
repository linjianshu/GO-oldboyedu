package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 5; i++ {
		go calculate(i, jobChan, resultChan)
	}
	for i := 0; i < 10; i++ {
		generate(int64(i), jobChan)
	}
	close(jobChan)

	for i := 0; i < 5; i++ {
		fmt.Println(<-resultChan)
	}
}

var (
	jobChan    = make(chan int, 100)
	resultChan = make(chan int, 100)
)

func calculate(id int, jobChan <-chan int, resultChan chan<- int) {
	for intn := range jobChan {
		fmt.Printf("goroutine %v start ,calcuting %v\n", id, intn)
		k := intn / 1000
		b := (intn - k*1000) / 100
		s := (intn - k*1000 - b*100) / 10
		g := intn - k*1000 - b*100 - s*10
		resultChan <- k + b + s + g
		fmt.Printf("goroutine %v end   ,calcuted  %v\n", id, intn)
	}
}

func generate(i int64, jobChan chan<- int) {
	rand.Seed(i)
	intn := rand.Intn(1000)
	jobChan <- intn
}
