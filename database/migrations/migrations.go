package migrations

import (
	"github.com/alefejonatha/Http/restapi/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) { //Passando os Models
	db.AutoMigrate(models.Book{})
}
