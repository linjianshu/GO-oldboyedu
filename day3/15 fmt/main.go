package main

import "fmt"

func main() {
	fmt.Print("hello world")
	fmt.Println("hello world")
	fmt.Printf("%p", "helloworld")

	//%d 十进制
	//%v 值
	//%o 八进制
	//%x 十六进制
	//%T 类型
	//%s 字符串
	//%p 指针
	//%b 二进制
	//%c 字符s
	//%f 浮点数
	//%t 布尔值
	fmt.Println()
	var m map[string]int
	m = make(map[string]int)
	m["ljs"] = 98
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)

	fmt.Printf("%q\n", 65)
	printfPersentage(98)
	fmt.Printf("%b\n", 5.6)

	n := 12.34
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.2f\n", n)
	fmt.Printf("%9.f\n", n)

	s := "小王子"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	//一共5个 只留2个
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)

	var s1 string
	fmt.Scan(&s1)
	fmt.Println(s1)

	var (
		name  string
		age   int
		class string
	)
	//fmt.Scanf("%s %d %s\n",&name,&age,&class)
	fmt.Printf("%s %d %s\n", name, age, class)
	fmt.Scanln(&name, &age, &class)
	fmt.Printf("%s %d %s\n", name, age, class)

}

func printfPersentage(a int) {
	fmt.Printf("%d%%\n", a)
}
