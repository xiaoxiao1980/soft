package router

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

// InitUserRouter 基础路由
func InitTokenRouter(Router *gin.RouterGroup) {
	TokenRouter := Router.Group("token")
	{
		TokenRouter.POST("", api.Login)
		// TokenRouter.DELETE("", v1.Logout)
	}
}
