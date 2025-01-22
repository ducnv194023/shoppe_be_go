package initialize

import (
	"database/sql"
	"fmt"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	
	"github.com/ducnv194023/shoppe_be_go/global"
	"github.com/ducnv194023/shoppe_be_go/internal/database/sqlc"
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

	// Create database connection
	db, err := sql.Open("mysql", s)
	checkErrorPanic(err, "Failed to connect to database")

	// Configure connection pool
	db.SetMaxIdleConns(m.MaxIdleConns)
	db.SetMaxOpenConns(m.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)

	// Run migrations
	runMigrations(db)

	// Initialize SQLC queries
	global.MDb = db
	global.Queries = database.New(db)
	
	global.Logger.Info("Database connection and queries initialized successfully!")
}

func runMigrations(db *sql.DB) {
	err := goose.SetDialect("mysql")
	if err != nil {
		global.Logger.Error("Failed to set database dialect", zap.Error(err))
		panic(err)
	}

	err = goose.Up(db, "internal/migrations")
	if err != nil {
		global.Logger.Error("Failed to run migrations", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Database migrations completed successfully!")
}