package database

import (
	"github.com/acger/user-svc/model"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
