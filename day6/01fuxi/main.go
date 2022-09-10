package main

import "fmt"

//类型断言
func main() {
	var a interface{}
	//fmt.Println("请输入一个值")
	//fmt.Scanf("%v\n",&a)
	//i,ok := a.(int)
	//if ok {
	//	fmt.Println("this is a int")
	//	fmt.Println(i)
	//}
	//fmt.Printf("%T\n",a)

	a = 100
	i, ok := a.(int)
	if ok {
		fmt.Println("this is a int ")
		fmt.Printf("%v\n", i)
	}

	switch i := a.(type) {
	case int:
		fmt.Println("this is a int")
		fmt.Println(i)
	}

}
