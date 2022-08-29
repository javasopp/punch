package settings

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ReadYamlByViper() {
	appConfig := &AppConfig{}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")
	viper.ReadInConfig()
	viper.Unmarshal(&appConfig)
	log.Info("我是viper读取到的值: ", appConfig.Mysql.Port)
}
