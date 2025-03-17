package routes

import (
	"order-service/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	//api version 1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "success",
				"message": "Order service API version 1",
			})
		})

		postGroupV1 := v1.Group("/orders")
		{
			postGroupV1.GET("/", handlers.GetOrders)
			postGroupV1.POST("/", handlers.CreateOrder)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Route not found",
		})
	})

}
