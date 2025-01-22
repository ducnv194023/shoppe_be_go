package global

import (
	"database/sql"

	"github.com/ducnv194023/shoppe_be_go/internal/database/sqlc"
	"github.com/ducnv194023/shoppe_be_go/pkg/logger"
	"github.com/ducnv194023/shoppe_be_go/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MDb *sql.DB
	Queries *database.Queries
	Rdb *redis.Client
)
