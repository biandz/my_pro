package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	topic = "user_event_dev"
	group = "user_event_dev_1"
	host  = "localhost:9092"
)

func getAddr() []string {
	return strings.Split(host, ",")
}
func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

// handler，核心的消费者业务实现
type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(s sarama.ConsumerGroupSession) error {
	fmt.Println("set up ....") // 当连接完毕的时候会通知这个，start
	return nil
}
func (exampleConsumerGroupHandler) Cleanup(s sarama.ConsumerGroupSession) error {
	fmt.Println("Cleanup") // end，当这一次消费完毕，会通知，这里最好commit
	return nil
}
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error { // consume
	for msg := range claim.Messages() { // 接受topic消息
		fmt.Println("接收消息：", string(msg.Value))
		fmt.Printf("[Consumer] Message topic:%q partition:%d offset:%d add:%d\n", msg.Topic, msg.Partition, msg.Offset, claim.HighWaterMarkOffset()-msg.Offset)
		sess.MarkMessage(msg, "") // 必须设置这个，不然你的偏移量无法提交。
	}
	return nil
}

func main() {
	go func() {
		http.ListenAndServe(":8888", http.DefaultServeMux) // pprof
	}()
	sarama.Logger = log.New(os.Stderr, "[SARAMA] ", log.LstdFlags) // 可以使用自定义日志存储，全局
	wg := sync.WaitGroup{}
	wg.Add(2)
	producer(&wg) // 生产者
	time.Sleep(10 * time.Second)
	fmt.Println("exec consumer")
	consumer(&wg) // 消费者
	wg.Wait()
}

func consumer(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		client, err := sarama.NewConsumerGroup(getAddr(), group, newKafkaConfig()) // broker_ip，消费者组(broker记录偏移量)，kafka 配置设置
		panicError(err)
		for { // for循环的目的是因为存在重平衡，他会重新启动
			handler := new(exampleConsumerGroupHandler)                    // 必须传递一个handler
			err = client.Consume(context.TODO(), []string{topic}, handler) // consume 操作，死循环。exampleConsumerGroupHandler的ConsumeClaim不允许退出，也就是操作到完毕。
			panicError(err)
			fmt.Println("re  balance")
		}
	}()
}

func newKafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.ClientID = "sarama_demo"                  //
	config.Version = sarama.V0_11_0_1                // kafka server的版本号
	config.Producer.Return.Successes = true          // sync必须设置这个
	config.Producer.RequiredAcks = sarama.WaitForAll // 也就是等待foolower同步，才会返回
	config.Producer.Return.Errors = true
	config.Consumer.Return.Errors = true
	config.Metadata.Full = false                                           // 不用拉取全部的信息
	config.Consumer.Offsets.AutoCommit.Enable = true                       // 自动提交偏移量，默认开启，说时候，我没找到手动提交。
	config.Consumer.Offsets.AutoCommit.Interval = time.Second              // 这个看业务需求，commit提交频率，不然容易down机后造成重复消费。
	config.Consumer.Offsets.Initial = sarama.OffsetOldest                  // 从最开始的地方消费，业务中看有没有需求，新业务重跑topic。
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // rb策略，默认就是range
	return config
}

func producer(wg *sync.WaitGroup) {
	go func() {
		config := newKafkaConfig()
		defer wg.Done()
		producer, err := sarama.NewSyncProducer(getAddr(), config) // producer，就很简单了
		panicError(err)
		buffer := bytes.Buffer{}
		for {
			buffer.Reset()
			//time.Sleep(time.Millisecond * 100)
			buffer.WriteString(fmt.Sprintf("curent: %v", time.Now().Format("2006-01-02 15:04:05")))
			partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
				Topic: topic,                              // 需要指定topic
				Value: sarama.ByteEncoder(buffer.Bytes()), // value，对于kafka来说不推荐传递key，因为容易造成分区不均匀。
			})
			panicError(err)
			fmt.Fprintf(os.Stdout, "[Producer] partition: %v, offset: %v, topic: %v\n", partition, offset, topic)
			time.Sleep(2 * time.Second)
		}
	}()
}
