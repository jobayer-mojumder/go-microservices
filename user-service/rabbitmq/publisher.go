package rabbitmq

import (
	"encoding/json"
	"log"
	"user-service/models"

	"github.com/streadway/amqp"
)

func PublishUserCreated(user models.User) {
	conn, _ := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	ch, _ := conn.Channel()
	defer conn.Close()
	defer ch.Close()

	body, _ := json.Marshal(user)
	err := ch.Publish("", "user_created", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		log.Println("Failed to publish message:", err)
	} else {
		log.Println("User event published:", user.ID)
	}
}
