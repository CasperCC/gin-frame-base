package route

import (
	"gin-frame-base/app/api"
	"gin-frame-base/app/middleware"
	"github.com/gin-gonic/gin"
)

func genUserRouter(r *gin.RouterGroup) {
	r.Use(middleware.CorsMiddleware)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "user-pong"})
	})
	r.GET("/test", api.UserApi.GetUserDetail)
}
