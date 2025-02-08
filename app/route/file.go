package route

import (
	"gin-frame-base/app/api"
	"gin-frame-base/app/middleware"
	"github.com/gin-gonic/gin"
)

func genFileRouter(r *gin.RouterGroup) {
	fileApi := api.InitializeFileApi()
	r.Use(middleware.CorsMiddleware, middleware.JwtAuth)

	r.POST("/upload", fileApi.Upload)
}
