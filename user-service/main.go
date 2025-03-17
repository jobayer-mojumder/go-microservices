package main

import (
	"log"
	"user-service/config"
	"user-service/database"
	"user-service/database/migrations"
	"user-service/routes"
	"user-service/utils"

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

	port := utils.GetEnv("PORT", "8082")

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":" + port)

	// print the port the server is running on to the console
	println("Server running on port: " + port)
}
