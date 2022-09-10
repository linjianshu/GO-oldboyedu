package main

import (
	"fmt"
	"github.com/nxadm/tail"
	"time"
)

func main() {
	//tail用法
	fileName := "./my.log"
	config := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的那个地方开始读
		ReOpen:      true,                                 //重新打开
		MustExist:   false,                                //文件不存在不报错
		Poll:        true,
		Pipe:        false,
		Follow:      true, //是否跟随
		MaxLineSize: 0,
		RateLimiter: nil,
		Logger:      nil,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed , err: ", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen , filename :%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg: ", line.Text)
	}

}
