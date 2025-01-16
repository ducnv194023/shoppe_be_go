package setting

type Config struct {
	ServerSetting ServerSetting `mapstructure:"server"`
	MysqlSetting MysqlSetting `mapstructure:"mysql"`
	LoggerSetting LoggerSetting `mapstructure:"logger"`
	RedisSetting RedisSetting `mapstructure:"redis"`
}

type MysqlSetting struct {
	Host         	string `mapstructure:"host"`
	Port     		int    `mapstructure:"port"`
	Username     	string `mapstructure:"username"`
	Password 		string `mapstructure:"password"`
	Dbname   		string `mapstructure:"dbname"`
	MaxIdleConns 	int    `mapstructure:"maxIdleConns"`
	MaxOpenConns 	int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"maxSize"`
	MaxBackups int `mapstructure:"maxBackups"`
	MaxAge int `mapstructure:"maxAge"`
	Compress bool `mapstructure:"compress"`
}

type RedisSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB int `mapstructure:"db"`
	PoolSize int `mapstructure:"poolSize"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}