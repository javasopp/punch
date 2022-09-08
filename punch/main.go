package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"punch/routers"
	"punch/settings"
)

func main() {
	r := gin.Default()
	routers.InitRouters(r)
	log.Info("启动成功。。。")
	log.Info(settings.Config.Server.Port)
	r.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务

}
