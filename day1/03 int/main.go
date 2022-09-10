package main

import "fmt"

func main() {
	i1 := 10
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1) //把十进制转成8进制
	fmt.Printf("%b\n", i1) //把十进制转成2进制
	fmt.Printf("%x\n", i1) //把十进制转成16进制

	//八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//十六进制
	i3 := 0x123
	fmt.Printf("%d\n", i3)

	fmt.Printf("%T\n", i1)

	//声明一个int8类型的 要明确指定类型，都则就是int类型
	i4 := int8(9)
	fmt.Println(i4)
}
