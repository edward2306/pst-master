package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pisatoo/pst-master/models"
)

func Migrate(db *gorm.DB) (err error) {
	db.AutoMigrate(models.LXD{})
	db.AutoMigrate(models.LXC{})

	return
}
