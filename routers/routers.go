package routers

import (
	"BlueBell/controller"
	"BlueBell/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	engine := gin.New()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册业务路由
	engine.POST("/signup", controller.SignUpHandler)

	//登录业务路由
	engine.POST("/login", controller.LoginHandler)


	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return engine
}
