package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func write() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("err cause: %v\n", err)
		return
	}

	//write
	fileObj.Write([]byte{97, 98, 99})
	fileObj.Write([]byte("this is a b c "))
	fileObj.WriteString("hello world!")
	defer fileObj.Close()
}

func writeByBufIo() {
	file, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("err cause : %v\n", err)
		return
	}

	writer := bufio.NewWriter(file)
	writer.WriteString("comeon baby!") //bufio是做了一个缓存
	writer.Flush()
	defer file.Close()
}

func writeByIoutil() {
	str := "hello 北京"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("error cause : %v\n", err)
		return
	}
}
func main() {
	//write()
	//writeByBufIo()
	writeByIoutil()

}
