package main

import (
	"io/ioutil"
)

func main() {
	copyFile("./xxcopy.txt", "./xx.txt")
}

func copyFile(dstName, srcName string) (written int64, err error) {
	//以读的方式打开文件
	//file, err := os.OpenFile(srcName, os.O_RDONLY, 0644)
	//if err != nil {
	//	return 0, err
	//}
	//reader:= bufio.NewReader(file)
	readFile, err := ioutil.ReadFile(srcName)
	if err != nil {
		return 0, err
	}

	err1 := ioutil.WriteFile(dstName, readFile, 0644)
	if err1 != nil {
		return 0, err1
	}

	return 1, nil

}
