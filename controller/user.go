package controller

import (
	"BlueBell/dao/mysql"
	"BlueBell/logic"
	"BlueBell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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
			ResponseError(context, CodeInvalidParam)
		}

		ResponseErrorWithMsg(context, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) // 翻译错误
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
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(context, CodeUserExist)
			return
		}
		ResponseError(context, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(context, nil)

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
			ResponseError(context, CodeInvalidParam)
		}

		ResponseErrorWithMsg(context, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) // 翻译错误

		return
	}
	// 业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))

		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(context, CodeUserNotExist)
			return
		}
		ResponseError(context, CodeInvalidPassword)
		return
	}

	// 返回响应
	ResponseSuccess(context, token)
}
