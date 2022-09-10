package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	//tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要 leader 和 follower 都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个 partitioner
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel 返回
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test blog")
	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err : ", err)
		return
	}
	defer client.Close()
	//发送消息
	pid, offSet, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed , err : ", err)
		return
	}
	fmt.Printf("pid:%v offSet:%v\n", pid, offSet)

}
