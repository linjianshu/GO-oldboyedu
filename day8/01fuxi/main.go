package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendNum(ch chan<- int) {
	for {
		intn := rand.Intn(10)
		time.Sleep(5 * time.Second)
		ch <- intn
	}
}
func main() {
	//channel
	var ch1 = make(chan int, 1)
	//ch1<-100  //把一个值发送到通道中
	//<-ch1	//把通道中把100取出来
	go sendNum(ch1)
	for {
		x, ok := <-ch1 //阻塞了
		if !ok {
			fmt.Println("!ok")
		}
		fmt.Println(x)
	}
}
