package router

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

func InitDataRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("data")
	{
		ConfigRouter.GET("field", api.GetField)
		ConfigRouter.GET("file", api.GetFile)
		ConfigRouter.GET("version", api.GetVersion)
		ConfigRouter.GET("record", api.GetRecord)
		ConfigRouter.GET("table", api.GetTable)
	}
}
