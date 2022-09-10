package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notifyCh = make(chan struct{}, 5)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		jobs <- i
	}

	close(jobs)

	go func() {
		//如果换成forrange
		for result := range results {
			fmt.Println(result)
		}
	}()

	wg.Wait()
	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %v is starting job %v\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("worker %v is ending job %v\n", id, job)
		results <- job * 2
		notifyCh <- struct{}{}

		wg.Done()
	}
}

func f1() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		jobs <- i
	}
	close(jobs)

	//for i := 0; i < 5; i++ {
	//	<-results
	//}

	go func() {
		//如果换成forrange
		for result := range results {
			fmt.Println(result)
		}
	}()

	wg.Wait()
	close(results)
}
