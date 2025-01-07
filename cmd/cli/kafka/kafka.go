package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19092"
	kafkaTopic = "user_topic_0001"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, // []string{"localhost01", "localhost"},
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset, // when a user starts, it must get the value first before the user
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// stock transaction
func newStock(msg, typeMessage string) *StockInfo {

	s := StockInfo{}
	s.Message = msg
	s.Type = typeMessage
	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)
	// create message
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}
	// write message by producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action Successfully",
	})
}

// consumer looking for buying ATC
func RegisterConsumeraATC(id int) {
	// group consume
	kafkaGroupId := "consumer-group-"
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer (%d) Hong Phien ATC::\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer (%d) error: %v", id, err)

		}
		fmt.Printf("Consumer (%d), hong topic:%v, partition:%v, offset:%v, time:%d %s = %s\n", id, m.Topic, m.Partition, m.Offset,
			m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	// regist hong
	go RegisterConsumeraATC(1)
	go RegisterConsumeraATC(2)

	r.Run(":8999")
}
