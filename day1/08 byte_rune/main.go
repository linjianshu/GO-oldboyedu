package main

import "fmt"

func main() {
	s := "Hello 沙河"
	n := len(s)
	fmt.Println(n)

	//这样会有问题
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)\n", s[i], s[i])
	}

	for _, i := range s {
		fmt.Printf("%v(%c)\n", i, i)
	}

	s1 := "big"
	s1 = "pig"
	fmt.Println(s1)

	s2 := "白萝卜"
	s2 = "红萝卜"
	fmt.Println(s2)

	//字符串修改
	s21 := "白萝卜"
	rune1 := []rune(s21)
	rune1[0] = '红'
	fmt.Println(string(rune1))

	//类型转换
	n1 := 10
	var float1 float32
	float1 = float32(n1)
	fmt.Println(float1)
	fmt.Printf("%T\n", float1)

}
