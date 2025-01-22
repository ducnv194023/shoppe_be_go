package initialize

import validator "github.com/ducnv194023/shoppe_be_go/internal/middleware"



func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
	validator.InitValidator()

	r := InitRouter()

	r.Run(":8080")
}
