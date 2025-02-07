package route

import "github.com/gin-gonic/gin"

type AppRouter struct{}

func (*AppRouter) AddRoute(e *gin.Engine) {
	genTestRouter(e.Group("/test"))
	genAdminRouter(e.Group("/admin"))
	genUserRouter(e.Group("/user"))
}

func New() *AppRouter {
	return &AppRouter{}
}
