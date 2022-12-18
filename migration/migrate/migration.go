package migrate

import (
	"github.com/mohnaofal/go-clean-architecture/app/models"
	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
	)
}
