package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)
var lock = sync.Mutex{}

func main() {
	//go内置的map不是并发安全的
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			set(key, i)
			fmt.Printf("k=%v v=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func get(key string) int {
	return m[key]
}
func set(key string, value int) {
	//lock.Lock()
	m[key] = value
	//lock.Unlock()
}
