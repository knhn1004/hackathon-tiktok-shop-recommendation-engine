package kafka

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"sync"

	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/config"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

var (
	articleWriter *kafka.Writer
	productWriter *kafka.Writer
	once          sync.Once
)

type KafkaMessage struct {
	Topic string
	Key   []byte
	Value interface{}
}

var kafkaChannel chan KafkaMessage

func init() {
	kafkaChannel = make(chan KafkaMessage, 1000) // Buffer size of 1000
	go kafkaWorker()
}

func initKafkaWriters() {
	mechanism, err := scram.Mechanism(scram.SHA256, config.KafkaUsername, config.KafkaPassword)
	if err != nil {
		panic(err)
	}

	articleWriter = &kafka.Writer{
		Addr:  kafka.TCP(config.KafkaUrl),
		Topic: config.KafkaArticleInteractionTopic,
		Transport: &kafka.Transport{
			SASL: mechanism,
			TLS:  &tls.Config{},
		},
	}

	productWriter = &kafka.Writer{
		Addr:  kafka.TCP(config.KafkaUrl),
		Topic: config.KafkaProductInteractionTopic,
		Transport: &kafka.Transport{
			SASL: mechanism,
			TLS:  &tls.Config{},
		},
	}
}

func GetArticleWriter() *kafka.Writer {
	once.Do(initKafkaWriters)
	return articleWriter
}

func GetProductWriter() *kafka.Writer {
	once.Do(initKafkaWriters)
	return productWriter
}

func WriteArticleInteraction(ctx context.Context, key []byte, value interface{}) {
	select {
	case kafkaChannel <- KafkaMessage{Topic: config.KafkaArticleInteractionTopic, Key: key, Value: value}:
		// Message sent to channel
	default:
		// Channel is full, log an error or implement a fallback strategy
		log.Printf("Kafka channel is full, article interaction message dropped")
	}
}

func WriteProductInteraction(ctx context.Context, key []byte, value interface{}) {
	select {
	case kafkaChannel <- KafkaMessage{Topic: config.KafkaProductInteractionTopic, Key: key, Value: value}:
		// Message sent to channel
	default:
		// Channel is full, log an error or implement a fallback strategy
		log.Printf("Kafka channel is full, product interaction message dropped")
	}
}

func kafkaWorker() {
	for msg := range kafkaChannel {
		valueJSON, err := json.Marshal(msg.Value)
		if err != nil {
			log.Printf("Error marshaling Kafka message: %v", err)
			continue
		}

		var writeErr error
		if msg.Topic == config.KafkaArticleInteractionTopic {
			writeErr = writeArticleInteractionToKafka(context.Background(), msg.Key, valueJSON)
		} else if msg.Topic == config.KafkaProductInteractionTopic {
			writeErr = writeProductInteractionToKafka(context.Background(), msg.Key, valueJSON)
		}

		if writeErr != nil {
			log.Printf("Error writing to Kafka: %v", writeErr)
			// Implement retry logic here if needed
		}
	}
}

func writeArticleInteractionToKafka(ctx context.Context, key, value []byte) error {
	return GetArticleWriter().WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

func writeProductInteractionToKafka(ctx context.Context, key, value []byte) error {
	return GetProductWriter().WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

func CloseWriters() error {
	if articleWriter != nil {
		if err := articleWriter.Close(); err != nil {
			return err
		}
	}
	if productWriter != nil {
		if err := productWriter.Close(); err != nil {
			return err
		}
	}
	return nil
}