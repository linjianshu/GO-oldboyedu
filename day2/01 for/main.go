package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//i==5时就跳出循环
		if i == 5 {
			fmt.Println(i)
		}
	}
	fmt.Println("over!")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //继续下一次循环
		}
		fmt.Println(i)
	}
	fmt.Println("over1!")

}
