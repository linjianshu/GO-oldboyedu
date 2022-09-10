package main

import (
	"bufio"
	"fmt"
	"net"
	"src/code.oldboyedu.com/day8/11nianbao_jiejue/protocol"
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
				reader := bufio.NewReader(conn)
				decode, err := protocol.Decode(reader)
				if err != nil {
					return
				}
				fmt.Println("received data : ", decode)
			}

		}(accept)

	}
}
