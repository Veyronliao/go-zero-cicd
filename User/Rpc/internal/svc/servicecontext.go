package svc

import (
	models "Bolog/User/Model"
	"Bolog/User/Rpc/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     models.DB,
	}
}
