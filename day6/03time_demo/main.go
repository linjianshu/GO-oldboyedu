package main

import (
	"fmt"
	"time"
)

func main() {
	f2()
}

func f1() {
	now := time.Now()
	fmt.Printf("%v \n", now)
	fmt.Printf("%v \n", now.Year())
	fmt.Printf("%v \n", now.Month())
	fmt.Printf("%v \n", now.Day())
	fmt.Printf("%v \n", now.Hour())
	fmt.Printf("%v \n", now.Minute())
	fmt.Printf("%v \n", now.Second())
	fmt.Println(now.Date())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	//time.Unix()
	unix := time.Unix(now.Unix(), 0)
	fmt.Println(unix)
	//时间间隔
	fmt.Println(time.Second)
	//now +24 hours
	fmt.Println(time.Now().Add(time.Hour * 24))
	//定时器
	//tick := time.Tick(time.Second)
	//for  i := range tick {
	//	fmt.Println(i)
	//}

	//格式化时间 把语言中时间对象 转换成字符串类型的时间
	//2021/08/10
	fmt.Println(time.Now().Format("2006/01/02"))
	fmt.Println(time.Now().Format("2006-1-2 15:04:05"))
	fmt.Println(time.Now().Format("2006-1-2 03:04:05"))
	fmt.Println(time.Now().Format("2006-1-2 15:04:05 PM"))
	fmt.Println(time.Now().Format("2006:01:02 15:04:05.000 PM"))

	//按照对应的格式 解析字符串类型的时间
	value, err := time.Parse("2006-01-02", "2019-05-20")
	if err != nil {
		fmt.Println(" err ", err)
		return
	}
	fmt.Println(value)

	fmt.Println(time.Now().Sub(time.Now().Add(-time.Hour)))

	fmt.Println("beginning")
	//sleep
	time.Sleep(time.Second * 2)
	fmt.Println("ending...")
}

func f2() {
	now := time.Now()
	//获取的是当前时区的时间
	fmt.Println(now)
	//按照东八区的时区和格式解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2021-08-11 21:33:05")
	//根据字符加载时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed , err :%v\n", err)
		return
	}
	//按照指定时区解析时间
	parseInLocation, err := time.ParseInLocation("2006-01-02 15:05:05", "2021-08-11 21:33:05", location)

	fmt.Println(time.Now().Sub(parseInLocation))

}
