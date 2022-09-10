package main

import (
	"fmt"
	"os"
)

func main() {
	//OpenFile()
	InsertFile()
}

func InsertFile() {
	//打开文件
	file, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("err cause : %v\n", file)
		return
	}
	defer file.Close()

	//读首两个字节
	var b = [2]byte{}
	n, _ := file.Read(b[:])

	//创建文件 写首两个字节
	openFile, _ := os.OpenFile("./sbinsert.txt", os.O_CREATE|os.O_WRONLY, 0644)
	openFile.Write(b[:n])
	defer openFile.Close()

	//尝试移动光标
	_, err1 := file.Seek(2, 0) //光标移动
	if err1 != nil {
		return
	}

	//尝试写要插入的数据
	openFile.Write([]byte{'c'})

	//尝试读光标下一个字节的数据
	var a [128]byte
	read, err2 := file.Read(a[:])
	if err2 != nil {
		return
	}

	fmt.Println(string(a[:read]))

	openFile.Write(a[:read])

	os.Rename("./sbinsert.txt", "./sb.txt")
	//writer := bufio.NewWriter(file)
	//writer.WriteString("c")
	//writer.Flush()

}
func OpenFile() {
	open, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Printf("err cause: %v\n", open)
		return
	}
	defer open.Close()
	var b [128]byte
	for {
		read, err1 := open.Read(b[:])
		if err1 != nil {
			fmt.Printf("read err cause: %v\n", err1)
			return
		}
		fmt.Println(string(b[:]))
		if read < 128 {
			return
		}
	}

}
