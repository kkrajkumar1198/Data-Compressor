package initializers

import (
	"log"

	"github.com/IBM/sarama"
)

var Worker sarama.Consumer
var Producer sarama.SyncProducer
var brokersUrl = []string{"localhost:9092", "localhost:9092"}

// Setting Up Kafka
func StartKafka() {
	var err error
	log.Println("Started Consumer")
	Worker, err = ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	log.Println("Started Producer")
	Producer, err = ConnectProducer(brokersUrl)
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected Producer")
	}
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func ConnectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
