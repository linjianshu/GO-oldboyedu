package main

import "fmt"

func main() {
	//基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//变种1
	i := 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	//变种2
	k := 5
	for k < 10 {
		fmt.Println(k)
		k++
	}

	////无限循环
	//for  {
	//	fmt.Println("无限循环")
	//}

	s := "Hello 你好中国"
	for i, v := range s {
		fmt.Printf("index:%d,value:%c\n", i, v)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
