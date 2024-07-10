package main

import (
	"basic-kafka/internal/kafka"
	"basic-kafka/internal/mongodb"
	"context"
	"log"
)

func main() {
	err := mongodb.InitMongoDB(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}
	kafka.RunConsumer()
}
