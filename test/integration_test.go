package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/kkrajkumar1198/Zocket/initializers"
	"google.golang.org/api/iterator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGCSClientIntegration(t *testing.T) {
	// Initialize the GCS client by calling GetGCSClient
	client, err := initializers.GetGCSClient()
	if err != nil {
		t.Fatalf("Failed to create GCS client: %v", err)
	}

	ctx := context.Background()

	// List all buckets in the project
	bucketIterator := client.Buckets(ctx, initializers.ProjectID)

	// Iterate over the buckets
	for {
		bucketAttrs, err := bucketIterator.Next()
		if err == nil {
		} else if err == iterator.Done {
			break
		} else {
			t.Fatalf("Error iterating over buckets: %v", err)
		}

		log.Printf("GCS Bucket Name: %s\n", bucketAttrs.Name)
	}

}

func TestDBConnectionIntegration(t *testing.T) {
	var db *gorm.DB
	var err error
	dsn := "provided the dsn string here"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	if db == nil {
		t.Fatal("Failed to connect to the test database.")
	}

	// Check if the DB connection is valid by executing a simple query
	var result int
	if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		t.Fatalf("Error executing test query: %v", err)
	}

	// Check if the result of the test query is as expected (should be 1)
	if result != 1 {
		t.Fatalf("Unexpected test query result: %d", result)
	}

	log.Println("TestDBConnectionIntegration: Database connection successful.")
}

func TestKafkaIntegration(t *testing.T) {

	// Start Kafka setup (Producer and Consumer)

	initializers.StartKafka()
	defer func() {

		// Close the Kafka Producer and Consumer at the end of the test
		if err := initializers.Producer.Close(); err != nil {
			t.Fatalf("Error closing Kafka Producer: %v", err)
		}
		if err := initializers.Worker.Close(); err != nil {
			t.Fatalf("Error closing Kafka Consumer: %v", err)
		}
	}()

	// Producing a test message
	testMessage := "Hello, Testing Kafka"
	producerMessage := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder(testMessage),
	}

	// Send the test message to Kafka
	_, _, err := initializers.Producer.SendMessage(producerMessage)
	if err != nil {
		t.Fatalf("Failed to send message to Kafka: %v", err)
	}

	// Consume the test message from Kafka
	consumer, err := initializers.Worker.ConsumePartition("test-topic", 0, sarama.OffsetOldest)
	if err != nil {
		t.Fatalf("Failed to create Kafka Consumer: %v", err)
	}

	// Wait for the message to be consumed
	select {
	case msg := <-consumer.Messages():
		receivedMessage := string(msg.Value)
		if receivedMessage != testMessage {
			t.Fatalf("Received unexpected message from Kafka. Expected: %s, Got: %s", testMessage, receivedMessage)
		}
	case err := <-consumer.Errors():
		t.Fatalf("Error while consuming Kafka message: %v", err)
	case <-time.After(10 * time.Second):
		t.Fatal("Timeout waiting for Kafka message")
	}
}
