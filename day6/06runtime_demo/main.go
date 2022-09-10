package main

import (
	"fmt"
	"runtime"
)

func main() {
	f1()
}

func f1() {
	caller, file, line, ok := runtime.Caller(0)
	if !ok {
		return
	}
	name := runtime.FuncForPC(caller).Name()
	fmt.Println(name)
	fmt.Println(file)
	fmt.Println(line)
}
