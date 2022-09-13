package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"punch/routers"
	"punch/settings"
	"strconv"
)

func main() {
	r := gin.Default()
	routers.InitRouters(r)
	log.Info("启动成功。。。")
	r.Run(fmt.Sprintf(":%s", strconv.Itoa(settings.Config.Server.Port)))
}
