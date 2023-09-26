package messagequeue

import (
	"log"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/kkrajkumar1198/Zocket/controllers"
	"github.com/kkrajkumar1198/Zocket/initializers"
	"github.com/kkrajkumar1198/Zocket/models"
)

func PushProductIDToQueue(topic string, id int) error {
	MessageValue := strconv.Itoa(id)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(MessageValue),
	}
	// Send the message to the topic
	partition, offset, err := initializers.Producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func NewProduct(c *gin.Context) {
	// Instantiate new Product struct
	var product models.Product
	c.Bind(&product)

	controllers.CreateProduct(product)

	log.Printf("From Producer: %d", product.ProductID)

	// Push the productID data to the Kafka queue
	PushProductIDToQueue("productID", product.ProductID)

	// Return Comment in JSON format
	c.JSON(200, gin.H{
		"success": true,
		"message": "Product pushed successfully",
	})
}
