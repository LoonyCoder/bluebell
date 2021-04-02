package middlewares

import (
	"BlueBell/controller"
	"BlueBell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization Bearer xxxxx.xxxxx.xxxxx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {

			controller.ResponseError(context, controller.CodeNeedLogin)

			context.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {

			controller.ResponseError(context, controller.CodeInvalidToken)

			context.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {

			controller.ResponseError(context, controller.CodeInvalidToken)

			context.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文context上
		context.Set(controller.CtxUserIDKey, mc.UserID)
		context.Next() // 后续的处理函数可以用过context.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}
