package route

import (
	"gin-frame-base/app/api"
	"gin-frame-base/app/middleware"
	"github.com/gin-gonic/gin"
)

// genUserRouter
// 用户路由
//
//	@param r
func genUserRouter(r *gin.RouterGroup) {
	userApi := api.InitializeUserApi()
	r.Use(middleware.CorsMiddleware)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "user-pong"})
	})
	r.POST("/login", userApi.Login)

	r.Use(middleware.JwtAuth)
	r.GET("/test", userApi.GetUserDetail)
}
