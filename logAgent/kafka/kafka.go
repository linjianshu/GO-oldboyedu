package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

//专门往kafka写日志的模块
var (
	client sarama.SyncProducer //声明一个全局的连接kafka的生产者client
)

//Init 初始化client
func Init(address []string) (err error) {
	config := sarama.NewConfig()
	//tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要 leader 和 follower 都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个 partitioner
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel 返回
	//连接kafka
	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Println("producer closed, err : ", err)
		return
	}
	return
}

func SendToKafka(topic, msg string) {
	//构造一个消息
	saramaMsg := &sarama.ProducerMessage{}
	saramaMsg.Topic = topic
	saramaMsg.Value = sarama.StringEncoder(msg)
	//发送消息
	pid, offSet, err := client.SendMessage(saramaMsg)
	if err != nil {
		fmt.Println("send msg failed , err : ", err)
		return
	}
	fmt.Printf("pid:%v offSet:%v\n", pid, offSet)
}
