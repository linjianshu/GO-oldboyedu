package main

import (
	"fmt"
	"net"
)

//tcp server端
func main() {
	//1.本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start server on 127.0.0.1:20000 failed , err : ", err)
		return
	}

	for {
		//2.等待别人来跟我连接
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("build connect failed , err : ", err)
			return
		}

		go func(conn net.Conn) {
			for {
				//3.与客户端通信
				var temp [128]byte
				read, err := conn.Read(temp[:])
				if err != nil {
					fmt.Println("attemp read failed , err : ", err)
					return
				}
				fmt.Println(string(temp[:read]))
			}
		}(accept)

	}

}
