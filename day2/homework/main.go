package main

import "fmt"

func main() {
	//单行注释
	/*
		多行注释
	*/

	//main函数是入口函数，他没有参数也没有返回值
	age := 19
	if age > 18 {
		fmt.Println("成年了！")
	} else if age > 7 {
		fmt.Println("上小学了！")
	} else {
		fmt.Println("最快乐的时光！")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	var i1 = 3
	for i1 < 10 {
		fmt.Println(i1)
		i1++
	}

	for {
		fmt.Println("hello world")
		break
	}

	for i, v := range "hello world" {
		fmt.Printf("index:%d\t,value:%c\n", i, v)
	}

	s := "hello world"
	for i, _ := range s {
		fmt.Printf("index:%d\t,value:%c\n", i, s[i])
	}

	//哑元变量，不想用到的都直接给_
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j <= 10-i; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	for i := 1; i < 10; i++ {
		for j := i; j > 0; j-- {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	//进制数
	s3 := 07
	s4 := 0xf
	fmt.Printf("%d\n", s3)
	fmt.Printf("%d\n", s4)

}
