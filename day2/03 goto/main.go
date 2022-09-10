package main

import "fmt"

func main() {
	flag := false
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				flag = true
				break
			}

			fmt.Printf("%d-%d\n", i, j)
		}
		if flag {
			break
		}
	}

	//for i := 0; i < 10; i++ {
	//	for j := 0; j < 10; j++ {
	//		if j == 2 {
	//			goto breakTag
	//		}
	//		fmt.Printf("%v-%v\n",i,j)
	//	}
	//}
	//return

	//breakTag:
	//	fmt.Println("结束for循环")
}
