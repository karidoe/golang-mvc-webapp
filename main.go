package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	ProductsController "golang-mvc-webapp/controllers/products"
	UsersController "golang-mvc-webapp/controllers/users"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"net/http"
)

func init() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:29092"})

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	go func() {
		for e := range producer.Events() {
			fmt.Println(e)
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "createOrder"
	for _, word := range []string{"{orderCreated:9000202}"} {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	producer.Flush(1 * 1000)
}

func main() {
	loadEnv()
	r := mux.NewRouter()

	ProductsController.BindRoutes(r)
	UsersController.BindRoutes(r)

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal("Serving error.", err)
	}
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
