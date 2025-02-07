package route

import (
	"gin-frame-base/app/api"
	"github.com/gin-gonic/gin"
)

func genTestRouter(r *gin.RouterGroup) {
	testApi := &api.TestApi{}
	r.GET("/buckets", testApi.Buckets)
	r.POST("/upload", testApi.Upload)
	r.GET("/file-list", testApi.FileList)
}
