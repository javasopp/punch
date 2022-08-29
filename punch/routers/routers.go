package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}
