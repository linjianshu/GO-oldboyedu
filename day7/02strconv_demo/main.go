package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "1000"
	parseInt, err := strconv.ParseInt(str, 10, 64) //10进制 int64
	if err != nil {
		return
	}
	fmt.Printf("%T %v\n", parseInt, parseInt)

	parseInt1, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return
	}
	fmt.Printf("%T %v\n", parseInt1, parseInt1)

	atoi, _ := strconv.Atoi("1000") //go语言继承c语言而来的 a是array 因为string底层实际上是array的byte数组 i是int
	fmt.Println(atoi)

	i := 97
	fmt.Println(string(i))  //a
	sprint := fmt.Sprint(i) //97

	//字符串中解析出bool值
	s2 := "true"
	parseBool, _ := strconv.ParseBool(s2)
	fmt.Println(parseBool)

	fmt.Println(strconv.ParseFloat("2.14", 32))

	fmt.Println(sprint)

}
