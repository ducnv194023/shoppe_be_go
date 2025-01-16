package initialize

import (
	"fmt"

	"github.com/ducnv194023/shoppe_be_go/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration file: %w", err))
	}

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration: %w", err))
	}
}
