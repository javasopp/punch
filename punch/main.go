package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"punch/routers"
	"punch/settings"
)

func init() {
	// TODO 读取配置文件
	settings.ReadYamlByViper()
}

func main() {
	r := gin.Default()
	routers.InitRouters(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	log.Info("启动成功。。。")
}
