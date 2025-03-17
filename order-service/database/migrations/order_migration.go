package migrations

import (
	"order-service/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateOrderTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "service_order_2025_03_18_001",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Order{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&models.Order{})
		},
	}
}
