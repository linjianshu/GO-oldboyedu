package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//tcp client
func main() {
	//1.与server建立连接
	dial, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:20000 failed , err : ", err)
		return
	}
	//2.发送数据
	//var write1 = make([]byte,100)
	//if len(os.Args)<2 {
	//	write1 = []byte("hello world")
	//}else {
	//	write1 = []byte(os.Args[1])
	//}

	for {
		reader := bufio.NewReader(os.Stdin)
		//dial.Write(write1)
		line, _ := reader.ReadString('\n')
		if string(line) == "exit" {
			break
		}
		dial.Write([]byte(line))
	}

	dial.Close()
}
