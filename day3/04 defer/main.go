package main

import "fmt"

func main() {
	deferDemo()
}

func deferDemo() {
	//defer多用于函数结束之前释放资源（文件句柄、数据库连接、socket链接）
	fmt.Println("start")
	defer fmt.Println("hhh")       //defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("heiheihei") //后进先出
	defer fmt.Println("xixixi")
	fmt.Println("end")
}
