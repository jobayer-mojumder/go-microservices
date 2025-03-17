package rabbitmq

import (
	"encoding/json"
	"log"
	"order-service/models"
	"sync"

	"github.com/streadway/amqp"
)

var (
	usersCache = make(map[uint]models.User)
	mutex      = sync.Mutex{}
)

func ListenForUserEvents() {
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

	// Declare the queue before consuming
	_, err = ch.QueueDeclare(
		"user_created", // Queue name
		true,           // Durable
		false,          // Auto-delete
		false,          // Exclusive
		false,          // No-wait
		nil,            // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare queue:", err)
	}

	msgs, err := ch.Consume(
		"user_created",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to consume from queue:", err)
	}

	for msg := range msgs {
		var user models.User
		json.Unmarshal(msg.Body, &user)

		mutex.Lock()
		usersCache[user.ID] = user
		mutex.Unlock()

		log.Println("User cached:", user.ID)
	}
}

func IsUserValid(userID uint) bool {
	mutex.Lock()
	defer mutex.Unlock()
	_, exists := usersCache[userID]
	return exists
}
