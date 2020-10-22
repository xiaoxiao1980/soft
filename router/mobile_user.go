package router

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

// InitUserRouter 基础路由
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("mobile")
	{
		UserRouter.POST("register", api.Register)
		UserRouter.POST("login", api.Login)
	}
}
