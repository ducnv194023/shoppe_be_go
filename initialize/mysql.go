package initialize

import (
	"fmt"
	"time"
	"go.uber.org/zap"
	"github.com/ducnv194023/shoppe_be_go/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/ducnv194023/shoppe_be_go/po"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(errString)
	}
}

func InitMysql() {
	m := global.Config.MysqlSetting
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	
	checkErrorPanic(err, "Failed to connect to database")
	global.Logger.Info("Connected to database!")
	global.MDb = db

	SetPool()
	migrateTables()
}

func SetPool() {
	sqlDB, err := global.MDb.DB()

	if err != nil {
		global.Logger.Error("Failed to set pool connection", zap.Error(err))
	}
	sqlDB.SetMaxIdleConns(global.Config.MysqlSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.Config.MysqlSetting.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(global.Config.MysqlSetting.ConnMaxLifetime))
}

func migrateTables() {
	err := global.MDb.AutoMigrate(
		&po.User{}, 
		&po.Role{},
	)

	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
	}
}

