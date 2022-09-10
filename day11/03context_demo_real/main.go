package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go f(ctx)
	//如何通知子goroutine退出
	time.Sleep(time.Second)
	//通知子goroutine退出
	cancel()
	wg.Wait()

}

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
LOOP:
	for {
		fmt.Println("ljs")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("jwt")
		time.Sleep(time.Millisecond * 250)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}
