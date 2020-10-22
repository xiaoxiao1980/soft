package router

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("sys")   // hao123.com/sys/register
	{
		BaseRouter.POST("register", api.SysRegister)
		BaseRouter.POST("login", api.SysLogin)
	}
	return BaseRouter
}
