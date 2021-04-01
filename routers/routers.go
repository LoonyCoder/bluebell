package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"BlueBell/logger"
)

func Setup() *gin.Engine {
	engine := gin.New()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return engine
}
