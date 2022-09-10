package main

import "fmt"

func main() {
	f1()
	f2()
}

func f1() {
	defer func() {
		i := recover()
		fmt.Printf("%v", i)
	}()
	panic("无法原谅！")
	fmt.Println("lalala")
}

func f2() {
	fmt.Println("lululu")
}
