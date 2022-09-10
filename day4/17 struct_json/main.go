package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//结构体与json
	//1.序列化 把go语言中的结构体变量 --> json格式的字符串
	//2.反序列化 把json格式的字符串  --> go语言中能够识别的结构体变量

	p1 := person{
		Name: "ljs",
		Age:  18,
	}

	//序列化
	v, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("marshal fail ")
		fmt.Println(err)
		fmt.Printf("%v  %T\n", err, err)
		return
	}
	fmt.Println(v)
	fmt.Printf("%v  %T\n", string(v), v)

	//反序列化  传指针进去
	var v1 person
	err1 := json.Unmarshal(v, &v1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(v1)
	fmt.Println(v1.Age)
	fmt.Println(v1.Name)

	s := `{"name":"ljs","age":18}`
	var v2 person //传指针进去
	json.Unmarshal([]byte(s), &v2)
	fmt.Printf("%#v\n", v2)

}

type person struct {
	Name string `json:"name" db:"dbname"`
	Age  int    `json:"age"`
}
