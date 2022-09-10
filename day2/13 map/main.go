package main

import "fmt"

func main() {
	var a map[string]int
	fmt.Println(a == nil)        //声明但是没有初始化 没有在内存中开辟空间
	a = make(map[string]int, 10) //要估算好该map的容量，避免在程序运行期间再动态扩容
	a["ljs"] = 18
	a["jwt"] = 3
	fmt.Println(a)
	fmt.Println(a["ljs"])

	fmt.Println(a["jww"]) //没有也不会报错 返回0
	//约定成俗 用ok来接收bool值
	value, ok := a["jww"]
	if !ok {
		fmt.Println("无此key")
	} else {
		fmt.Println(value)
	}

	for i, v := range a {
		fmt.Printf("index:%v value:%v\n", i, v)
	}

	//只遍历key
	for k := range a {
		fmt.Println(k)
	}
	//只遍历value
	for _, value := range a {
		fmt.Println(value)
	}

	delete(a, "ljs")
	fmt.Println(a)
	delete(a, "lllll") //删除不存在的不报错
	fmt.Println(a)
}
