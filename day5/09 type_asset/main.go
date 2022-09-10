package main

import (
	"fmt"
)

//类型断言
func main() {
	assert("100")
	assert(float32(32.1))
}

func assert(a interface{}) {
	fmt.Printf("%T %v\n", a, a)
	s, ok := a.(string) //类型断言
	if !ok {
		fmt.Println("error ")
		return
	}
	fmt.Println(s)

	switch i := a.(type) {
	case string:
		fmt.Printf("this is a string %T %v\n", i, i)
	case int:
		fmt.Printf("this is a int %T %v\n", i, i)
	case bool:
		fmt.Printf("this is a bool %T %v\n", i, i)
	case float64, float32:
		fmt.Printf("this is a float %T %v\n", i, i)

	}
}
