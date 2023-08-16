package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
	"time"
)

func main() {
	//建立链接
	conn, err := amqp.Dial("amqp://admin:123456@1.14.59.249:5672/")
	if err != nil {
		fmt.Println("建立链接失败：", err.Error())
		return
	}
	defer conn.Close()

	//打开一个通道
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("开启通道失败：", err.Error())
		return
	}
	defer ch.Close()

	//声明一个队列
	q, err := ch.QueueDeclare(
		"hello", // 队列名
		false,   // 是否持久化
		false,   // 是否自动删除
		false,   // 是否独占
		false,   // 是否阻塞等待
		nil,     // 参数
	)
	if err != nil {
		fmt.Println("声明一个队列失败：", err.Error())
		return
	}

	//发送消息
	go func() {
		var i = 0
		for i < 100 {
			body := "hello,world" + strconv.Itoa(i)
			err = ch.Publish(
				"",     // 交换机
				q.Name, // 队列名
				false,  // 强制持久化
				false,  // 立即发送
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			if err != nil {
				fmt.Println("发送消息失败：", err.Error())
				return
			}
			i++
		}
		fmt.Println("消息发送完毕！！！！")
	}()
	time.Sleep(30 * time.Second)
	//接收消息
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, amqp.Table{})
	if err != nil {
		fmt.Println("接收消息失败：", err.Error())
		return
	}

	go func() {
		for msg := range msgs {
			fmt.Println("接收到消息：", string(msg.Body))
		}
	}()

	select {}
}
