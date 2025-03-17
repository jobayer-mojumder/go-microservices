package repositories

import (
	"order-service/database"
	"order-service/models"
)

func CreateOrder(order *models.Order) error {
	err := database.DB.Create(&order).Error
	return err
}

func GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Find(&orders).Error
	return orders, err
}
