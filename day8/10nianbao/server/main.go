package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen port failed , err : ", err)
		return
	}
	defer listen.Close()

	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("connect accept failed , err :", err)
			return
		}

		go func(conn net.Conn) {
			defer conn.Close()
			for {
				var b [1024]byte
				read, err := conn.Read(b[:])
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("try to read failed , err : ", err)
					return
				}
				fmt.Println("received data : ", string(b[:read]))
			}

		}(accept)

	}
}
