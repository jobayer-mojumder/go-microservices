package main

import (
	"log"
	"order-service/config"
	"order-service/database"
	"order-service/database/migrations"
	"order-service/rabbitmq"
	"order-service/routes"
	"order-service/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()

	database.ConnectDB()

	// run all migrations
	err := migrations.RunMigrations(database.DB)
	if err != nil {
		log.Fatal(err)
	}

	// Start RabbitMQ consumer
	go rabbitmq.ListenForUserEvents()

	port := utils.GetEnv("PORT", "3001")

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":" + port)

	// print the port the server is running on to the console
	println("Server running on port: " + port)
}
