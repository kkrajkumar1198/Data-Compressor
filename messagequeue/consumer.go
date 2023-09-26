package messagequeue

import (
	"log"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/kkrajkumar1198/Zocket/cloudbucket"
	"github.com/kkrajkumar1198/Zocket/initializers"
	"github.com/kkrajkumar1198/Zocket/models"
)

func Start() {
	log.Println("Started Listening")
	topic := "productID"

	// Calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := initializers.Worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	/// Continuously consume messages using a for range loop
	for msg := range consumer.Messages() {

		var product models.Product
		stringValue := string(msg.Value)

		// Convert the string value to an integer
		product.ProductID, err = strconv.Atoi(stringValue)
		if err != nil {
			log.Printf("Error converting message to integer: %v", err)
			break
		}
		log.Printf("From Consumer - ProductID to fetch: %d", product.ProductID)

		result := initializers.DB.First(&product, product.ProductID)

		if result.Error != nil {
			log.Println("Cannot able to fetch the product from table :", result.Error)
			return
		}

		Images_names := product.ProductImages

		for _, image := range Images_names {
			log.Printf("Retrieved Images Name from DB: %s", image)
		}

		compressed_file_links, err := cloudbucket.DownloadAndCompressImages(Images_names)
		if err != nil {
			log.Printf("Failed to Download and Compressed the images:%s", err)
		} else {
			log.Println(compressed_file_links)
		}

		models.UpdateCompressedDataLocation(product.ProductID, compressed_file_links)
	}

}
