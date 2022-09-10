package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed...")
		return
	}

	//记得关闭文件
	defer fileObj.Close()
	var b = make([]byte, 128)
	for {
		n, err := fileObj.Read(b)
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Println("read from file failed , error")
			return
		}
		fmt.Println(n)
		fmt.Println(string(b))
		if n < 128 {
			return
		}
	}
}

//利用bufio这个包读取文件
func readFromFileByBufio() {
	fileObjFile, err := os.Open("./console.go")
	if err != nil {
		fmt.Printf("err, %v", err)
		return
	}
	defer fileObjFile.Close()

	reader := bufio.NewReader(fileObjFile)
	for {
		string, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed , err : %v", err)
			return
		}
		fmt.Print(string)
	}
}

func readFromFileByIoutil() {
	file, err := ioutil.ReadFile("./console.go")
	if err != nil {
		fmt.Printf("err , cause: %v\n", err)
	}
	fmt.Println(string(file))
}

//打开文件
func main() {
	//readFromFile1()
	//readFromFileByBufio()
	readFromFileByIoutil()
}
