package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	//defer会先把预定的值先算出来等着最后执行函数
	//defer calc("1",1,calc("10",1,2))
	//输出 "10" 1 2 3
	//defer calc("1",1,3)
	//a=0
	//defer calc("2",1,calc("20",0,2))
	//输出 "20" 0 2 2
	//defer calc("2",0,2)
	//b=1
	//程序退出
	//执行 defer calc("2",0,2)
	//输出 "2" 0 2 2
	//执行 defer calc("1",1,3)
	//输出 "1" 1 3 4
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
