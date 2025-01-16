package initialize

import (
	"github.com/ducnv194023/shoppe_be_go/global"
	"github.com/ducnv194023/shoppe_be_go/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LoggerSetting)
}
