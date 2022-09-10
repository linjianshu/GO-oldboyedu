package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var lock = sync.Mutex{}
var rwLock = sync.RWMutex{}
var wg = sync.WaitGroup{}

func main() {
	now := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(now))
}

func read() {
	rwLock.RLock()
	//lock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	wg.Done()
	rwLock.RUnlock()
}

func write() {
	rwLock.Lock()
	//lock.Lock()
	x++
	time.Sleep(time.Millisecond * 10)
	//lock.Unlock()
	wg.Done()
	rwLock.Unlock()
}
