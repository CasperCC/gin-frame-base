package route

import (
	"fmt"
	"gin-frame-base/app/model"
	"gin-frame-base/internal/global"
	"github.com/gin-gonic/gin"
)

func genUserRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "user-pong"})
	})
	r.GET("/test", func(c *gin.Context) {
		var user model.User
		global.Db.First(&user)
		fmt.Println(user)
		c.JSON(200, gin.H{"data": user})
	})
}
