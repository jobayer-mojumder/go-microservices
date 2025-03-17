package handlers

import (
	"net/http"
	"order-service/http/requests"
	"order-service/models"
	"order-service/rabbitmq"
	"order-service/repositories"
	"order-service/utils"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	var orderRequest requests.CreateOrderRequest
	if err := c.ShouldBind(&orderRequest); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	order := models.Order{
		ID:     orderRequest.Id,
		Total:  orderRequest.Total,
		UserID: orderRequest.UserID,
	}

	if !rabbitmq.IsUserValid(order.UserID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := repositories.CreateOrder(&order); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, gin.H{"order": order})
}

func GetOrders(c *gin.Context) {
	orders, err := repositories.GetOrders()
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, gin.H{"orders": orders})
}
