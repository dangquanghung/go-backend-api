package global

import (
	"github.com/dangquanghung/go-ecommerce-backend-api/pkg/logger"
	"github.com/dangquanghung/go-ecommerce-backend-api/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	// Mdb *gorm.DB
)

/*
MySql
Kafka
....

*/
