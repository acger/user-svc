package svc

import (
	"github.com/acger/user-svc/database"
	"github.com/acger/user-svc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	DB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DB : database.NewMysql(&c),
	}
}
