package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)


// 全局 Writer 变量，避免每次请求都创建连接
var writer *kafka.Writer


// InitProducer 初始化 Kafka 生产者
// 在 main.go 启动时调用一次即可
func InitProducer() {
	writer = &kafka.Writer{
		
		Addr:     kafka.TCP("localhost:9092"), 	// Kafka 地址，如果是 Docker 启动，注意是 localhost:9092 还是 kafka:9092
		Topic:    "milktea-orders",				// 消息发送到的 Topic
		Balancer: &kafka.LeastBytes{},			// 负载均衡策略：LeastBytes 最少字节优先，有助于数据均匀

		// 可选：设置异步写入，提升响应速度（虽然 WriteMessages 本身会等待 Batch）
		// Async: true, 
	}
	fmt.Println("Kafka Producer initialized success")
}


// SendOrderMessage 发送订单消息
// orderID: 订单ID
// payload: 具体的订单数据（比如 JSON 字符串）
func SendOrderMessage(orderID uint, payload interface{}) error {
	// 将 payload 对象转为 JSON 字节
	value, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("json marshal error: %v", err)
	}

	// key 转为 string，用来保证同一个订单ID的消息可能会去往同一个分区（如果需要有序性）
	key := fmt.Sprintf("%d", orderID)

	// 发送消息
	// 使用 context.Background() 也可以，或者传入带超时的 ctx
	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),   // 消息键 (用于分区路由)
			Value: value,         // 消息体 (你的订单数据)
		},
	)

	if err != nil {
		log.Printf("Failed to write message: %v", err)
		return err
	}

	log.Printf("Message sent to Kafka | Key: %s | Value: %s", key, string(value))
	return nil
}


// CloseProducer 关闭连接
// 在 main.go 退出时调用
func CloseProducer() {
	if writer != nil {
		if err := writer.Close(); err != nil {
			log.Printf("Failed to close kafka writer: %v", err)
		} else {
			fmt.Println("Kafka Producer closed")
		}
	}
}