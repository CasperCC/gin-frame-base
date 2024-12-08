package route

import "github.com/gin-gonic/gin"

func genAdminRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "admin-pong"})
	})
}
