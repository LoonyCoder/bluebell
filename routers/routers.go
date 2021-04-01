package routers

import (
	"BlueBell/controller"
	"BlueBell/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	engine := gin.New()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册业务路由
	engine.POST("/signup", controller.SignUpHandler)
	engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return engine
}
