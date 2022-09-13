package routers

import (
	"github.com/gin-gonic/gin"
	"punch/common"
)

func InitRouters(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, common.Success())
	})

	r.POST("/login")

	r.Group("/plan")

	r.Group("/punch")

	r.Group("/login")

}
