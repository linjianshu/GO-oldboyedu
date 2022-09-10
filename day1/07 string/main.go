package main

import (
	fmt "fmt"
	"strings"
)

func main() {
	path := "\"C:\\Users\\Sweetie\\Desktop\\车间级MES\""
	fmt.Printf("%s\t", path)

	s := `
	世情薄
	人情恶
	雨送黄昏花易落
`
	fmt.Printf("%s\r", s)
	s3 := `C:\Users\Sweetie\Desktop\车间级MES`
	fmt.Printf("%s\n", s3)

	//字符串相关操作
	fmt.Printf("%d\n", len(s3))

	//字符串拼接
	name := "ljs"
	world := "shuaibi"
	describtion := name + world
	fmt.Printf("%v\n", describtion)
	describtion1 := fmt.Sprintf("%s%s", name, world)
	fmt.Printf("%s\n", describtion1)

	//分割
	s1 := strings.Split(s3, "\\")
	fmt.Println(s1)
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	//包含
	fmt.Println(strings.Contains(describtion, name))
	fmt.Printf(" '%s' Contains '%s' ? result:%v", describtion1, name, strings.Contains(describtion, name))
	fmt.Println()
	//前缀、后缀
	fmt.Println(strings.HasPrefix(describtion, "ljs"))
	fmt.Println(strings.HasSuffix(describtion, "shuaibi"))

	//索引 查找
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "b"))
	fmt.Println(strings.LastIndex(s4, "b"))

	//拼接
	var sJoin = strings.Join(s1, "+")
	fmt.Println(sJoin)
}
