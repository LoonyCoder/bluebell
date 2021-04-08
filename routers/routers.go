package routers

import (
	"BlueBell/controller"
	"BlueBell/logger"
	"BlueBell/middlewares"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	_ "BlueBell/docs" // 千万不要忘了导入把你上一步生成的docs

)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	engine := gin.New()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册swagger api相关路由
	engine.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	v1 := engine.Group("/api/v1")
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)

	//登录业务路由
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或分数获取贴子列表
		v1.GET("/posts2", controller.GetPostListHandler2)

		// 投票
		v1.POST("/vote", controller.PostVoteHandler)

	}

	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return engine
}
