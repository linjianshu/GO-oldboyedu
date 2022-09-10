package main

import (
	"fmt"
	"os"
)

//1.文件对象的类型
//2.获取文件对象的详细信息
func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		return
	}
	fmt.Printf("%T\n", fileObj)

	fileInfo, err := fileObj.Stat()
	if err != nil {
		return
	}

	println(fileInfo.Size())

}
