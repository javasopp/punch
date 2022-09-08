package settings

import (
	"github.com/spf13/viper"
)

var Config *AppConfig

func init() {
	Config = &AppConfig{}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")
	viper.ReadInConfig()
	viper.Unmarshal(&Config)
}
