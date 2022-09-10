package main

import (
	"fmt"
)

func main() {
	score := 68
	switch score {
	case 68:
		fmt.Println("及格")
	default:
		fmt.Println("未知")
	}

	//简化代码 作用域问题
	switch i := 3; i {
	case 1:
		fmt.Println("wumingzhi")
	case 2:
		fmt.Println("zhongzhi")
	case 3:
		fmt.Println("damuzhi")
	}

	//同时声明几种情况
	switch i := 10; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("this is 奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("this is 偶数")
	}

	score1 := 68
	switch {
	case score1 > 60 && score1 <= 100:
		fmt.Println("及格")
	case score1 < 60:
		fmt.Println("挂了呀")

	}

}
