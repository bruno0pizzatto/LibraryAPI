package migrations

import (
	"github.com/bruno0pizzatto/LibraryAPI/models"
	"gorm.io/gorm"
)

func RunMMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.User{})
}
