package main

import (
	"flag"
	"fmt"
	"time"
)

//flag 获取命令行参数
func main() {
	//创建一个标志位参数flag
	name := flag.String("name", "ljs", "请输入名字")
	age := flag.Int("age", 100, "请输入年龄")
	married := flag.Bool("married", false, "结婚了么")
	cTime := flag.Duration("duration", time.Second, "有多快")
	//使用flag
	//flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)

	var name1 string
	flag.StringVar(&name1, "name1", "name1", "请输入name1")
	flag.Parse()
	fmt.Println(name1)

	fmt.Println(flag.NArg())  //返回除了规定的参数之外的命令行参数有几个
	fmt.Println(flag.NFlag()) //返回规定的flag命令行参数有几个
	fmt.Println(flag.Args())  //返回除了规定的参数之外的命令行参数具体是

}
