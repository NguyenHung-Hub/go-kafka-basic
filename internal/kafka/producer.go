package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"basic-kafka/configs"

	"github.com/segmentio/kafka-go"
)

// type Comment struct {
// 	Username  string
// 	Content   string
// 	CreatedAt time.Time
// }

func generateRandomUsername(seed int) string {
	return fmt.Sprintf("user_%d", seed)
}

func RunProducer() {
	topic := "comments"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", configs.KafkaAddress, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	for i := 0; i < 10000; i++ {
		comment := Comment{
			Username:  generateRandomUsername(i),
			Content:   "This is a comment",
			CreatedAt: time.Now(),
		}

		message := kafka.Message{
			Key:   []byte(comment.Username),
			Value: []byte(fmt.Sprintf("%s|%s|%s", comment.Username, comment.Content, comment.CreatedAt)),
		}

		_, err = conn.WriteMessages(message)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		log.Println("success to write messages ", i)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
