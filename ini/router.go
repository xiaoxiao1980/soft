package ini

import (
	_ "app/docs"
	"app/global"
	. "app/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Routers 路由
func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.Static("/assets", "./assets")

	// Router.StaticFS("/files", http.Dir("assets")) //直接进文件系统
	// Router.StaticFile("/favicon.ico", "./assets/0001.jpg")

	// swagger；注意：生产环境可以注释掉
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	APIGroup := Router.Group("")
	InitBaseRouter(APIGroup)
	InitTestRouter(APIGroup)   // 测试路由
	InitMongoRouter(APIGroup)  // mongo
	InitUserRouter(APIGroup)   // 用户路由
	InitFileRouter(APIGroup)   // 文件路由
	InitConfigRouter(APIGroup) //配置路由
	InitDataRouter(APIGroup)   // 档案数据组
	return Router
}

// RunWebServer 开启服务
func RunWebServer() {
	Router := Routers()
	s := &http.Server{
		Addr:           global.Config.Web.Addr,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
