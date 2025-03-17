package rabbitmq

import (
	"encoding/json"
	"log"
	"user-service/models"

	"github.com/streadway/amqp"
)

func PublishUserCreated(user models.User) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open channel:", err)
	}
	defer conn.Close()
	defer ch.Close()

	// Declare the queue before publishing
	_, err = ch.QueueDeclare(
		"user_created",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare queue:", err)
	}

	body, _ := json.Marshal(user)
	err = ch.Publish(
		"",
		"user_created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Failed to publish message:", err)
	} else {
		log.Println("User event published:", user.ID)
	}
}
