package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"github.com/ducnv194023/shoppe_be_go/pkg/logger"
	"github.com/ducnv194023/shoppe_be_go/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MDb *gorm.DB
	Rdb *redis.Client
)
