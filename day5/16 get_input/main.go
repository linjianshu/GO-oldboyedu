package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//useScan()
	ioscan()
}

func useScan() {
	fmt.Println("请输入内容！")
	var s string
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是 %v\n", s)
}

func ioscan() {
	fmt.Println("请输入内容！")
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println(s)
}
