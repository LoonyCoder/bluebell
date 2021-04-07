package controller

import (
	"BlueBell/logic"
	"BlueBell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 投票
func PostVoteHandler(context *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)

	if err := context.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(context, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体
		ResponseErrorWithMsg(context, CodeInvalidParam, errData)
		return
	}

	// 获取当前请求的用户的ID
	userId, err := GetCurrentUser(context)
	if err != nil {
		ResponseError(context, CodeNeedLogin)
		return
	}
	// 投票的业务逻辑
	if err := logic.VoteForPost(userId, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, nil)
}
