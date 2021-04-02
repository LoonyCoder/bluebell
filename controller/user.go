package controller

import (
	"BlueBell/logic"
	"BlueBell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(context *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := context.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			context.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		context.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
		})
		return
	}
	//// 手动对请求参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	// 请求参数有误，直接返回响应
	//	zap.L().Error("SignUp with invalid param")
	//	context.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	fmt.Println(p)
	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 3. 返回响应
	context.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func LoginHandler(context *gin.Context) {

	// 获取请求参数和参数校验
	p := new(models.ParamLogin)
	if err := context.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			context.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		context.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
		})
		return
	}
	// 业务逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		context.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}
	// 返回响应
	context.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
