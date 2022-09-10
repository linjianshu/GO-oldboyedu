package main

import "fmt"

//float
func main() {
	//maxFloat32 := math.MaxFloat32 最大值
	f1 := 1.23
	//默认go语言中的小数都是float64类型
	fmt.Printf("%T\n", f1)

	//显式声明float32类型
	f2 := float32(1.23)
	fmt.Printf("%T\n", f2)

	f1 = float64(f2) //不能隐式转换

}
