package tail

import (
	"fmt"
	"github.com/nxadm/tail"
)

var tails *tail.Tail

//Init 专门收集日志
func Init(address string) (err error) {
	//tail用法
	fileName := address
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
	tails, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed , err: ", err)
		return
	}
	return
}
func ReadLog() <-chan *tail.Line {
	return tails.Lines
}
