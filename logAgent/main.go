package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"src/code.oldboyedu.com/logAgent/kafka"
	tail "src/code.oldboyedu.com/logAgent/tail_log"
	"time"
)

type appConf struct {
	KafkaConf   `ini:"Kafka"`
	TailLogConf `ini:"Taillog"`
}
type KafkaConf struct {
	Address string `ini:"Address"`
	Topic   string `ini:"Topic"`
}

type TailLogConf struct {
	FileName string `ini:"Filename"`
}

var appCfg = new(appConf)

func main() {
	//0.加载配置文件 获取ip:端口 日志文件 发送的topic
	err := ini.MapTo(appCfg, "./config.ini")
	if err != nil {
		fmt.Println("config init failed, err : ", err)
		return
	}

	//1.初始化kafka连接
	fmt.Println(appCfg.KafkaConf.Address)
	fmt.Println(appCfg.KafkaConf.Topic)
	fmt.Println(appCfg.TailLogConf.FileName)
	err = kafka.Init([]string{appCfg.KafkaConf.Address})
	if err != nil {
		fmt.Println("init kafka failed, err : ", err)
		return
	}
	fmt.Println("init kafka success!")
	//2.打开日志文件准备收集日志
	err = tail.Init(appCfg.TailLogConf.FileName)
	if err != nil {
		fmt.Println("init taillog failed, err : ", err)
		return
	}
	fmt.Println("init tail success!")

	run()

}

func run() {
	//1.收集日志
	for {
		select {
		case line := <-tail.ReadLog():
			//2.发送给kafka
			kafka.SendToKafka(appCfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}
