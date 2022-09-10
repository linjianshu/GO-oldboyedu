package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	scoreMap := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	keys := make([]string, 200)
	for s, _ := range scoreMap {
		keys = append(keys, s)
	}

	sort.Strings(keys)
	for _, value := range keys {
		fmt.Println(value, scoreMap[value])
	}

}
