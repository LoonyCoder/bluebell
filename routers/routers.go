package routers

import (
	"BlueBell/controller"
	"BlueBell/logger"
	"BlueBell/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	engine := gin.New()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := engine.Group("/api/v1")
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)

	//登录业务路由
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())  // 应用JWT认证中间件

	{
		v1.GET("/community",controller.CommunityHandler)
		v1.GET("/community/:id",controller.CommunityDetailHandler)

		v1.POST("/post",controller.CreatePostHandler)
	}


	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return engine
}
