package router

import (
	"app/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// InitTestRouter 初始化测试路由
func InitTestRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("test").Use(middleware.JWTAuth())
	{
		// hello
		APIRouter.GET("hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		//APIRouter.GET("createTx", v1.CreateTx)
	}
}
