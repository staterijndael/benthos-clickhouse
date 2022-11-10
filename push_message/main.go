package main

import (
	"benthos/generated"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewSyncProducer(brokerList []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func main() {
	producer, err := NewSyncProducer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	accountPb := &schema.Person{
		TransactionId:  uuid.New().String(),
		FromUserId:     5,
		ToUserId:       5,
		CurrencySymbol: "BTC",
		Amount:         "50.0",
		AmountUsd:      "123123123.0",
	}

	marshalledProto, err := proto.Marshal(accountPb)
	if err != nil {
		panic(err)
	}

	msg := &sarama.ProducerMessage{
		Value: sarama.ByteEncoder(marshalledProto),
		Topic: "local_statistics",
	}

	producer.SendMessage(msg)
}
