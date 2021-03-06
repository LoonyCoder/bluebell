package routers

import (
	"BlueBell/controller"
	_ "BlueBell/docs" // 千万不要忘了导入把你上一步生成的docs
	"BlueBell/logger"
	"BlueBell/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	engine := gin.New()
	//engine.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册swagger api相关路由
	engine.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	engine.LoadHTMLFiles("./templates/index.html")
	engine.Static("/static", "./static")
	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	v1 := engine.Group("/api/v1")
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)

	//登录业务路由
	v1.POST("/login", controller.LoginHandler)

	// 根据时间或分数获取贴子列表
	v1.GET("/posts2", controller.GetPostListHandler2)
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/posts", controller.GetPostListHandler)
	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		v1.POST("/post", controller.CreatePostHandler)
		// 投票
		v1.POST("/vote", controller.PostVoteHandler)
	}

	// 注册pprof相关路由
	pprof.Register(engine)

	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return engine
}
