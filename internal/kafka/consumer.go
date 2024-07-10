package kafka

import (
	"context"
	"log"
	"strings"
	"time"

	"basic-kafka/configs"
	"basic-kafka/internal/mongodb"

	"github.com/segmentio/kafka-go"
)

func RunConsumer() {
	topic := "comments"
	partition := 0

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{configs.KafkaAddress},
		Topic:     topic,
		Partition: partition,
	})

	commentsChan := make(chan Comment, 10000) // Buffered channel to hold comments

	// Start a goroutine to handle batch inserts
	go handleBatchInserts(commentsChan)

	// Start multiple worker goroutines to handle batch inserts
	// numWorkers := 10
	// for i := 0; i < numWorkers; i++ {
	//     go handleBatchInserts(commentsChan)
	// }

	ctx := context.Background()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		parts := strings.Split(string(m.Value), "|")
		if len(parts) < 2 {
			log.Printf("Invalid message format: %v", string(m.Value))
			continue
		}

		comment := Comment{
			Username:  parts[0],
			Content:   parts[1],
			CreatedAt: time.Now(),
		}

		commentsChan <- comment
	}
}

func handleBatchInserts(commentsChan chan Comment) {
	collection := mongodb.GetMongoCollection()
	ctx := context.Background()
	batchSize := 100
	var comments []interface{}

	for comment := range commentsChan {
		comments = append(comments, comment)

		if len(comments) >= batchSize {
			_, err := collection.InsertMany(ctx, comments)
			if err != nil {
				log.Printf("Failed to insert comments: %v", err)
			} else {
				log.Printf("Inserted %d comments", len(comments))
			}
			comments = comments[:0] // Clear the slice for the next batch
		}
	}
}
