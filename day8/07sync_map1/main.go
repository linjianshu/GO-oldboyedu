package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.Map{}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			set(key, i)                              //必须使用sync.map内置的store方法设置键值对
			fmt.Printf("k=%v v=%v\n", key, get(key)) //必须使用sync.map内置的load方法根据key取值
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func set(key string, value int) {
	m.Store(key, value)
}

func get(key string) interface{} {
	value, _ := m.Load(key)
	return value
}
