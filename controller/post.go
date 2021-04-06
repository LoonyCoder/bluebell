package controller

import (
	"BlueBell/logic"
	"BlueBell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(context *gin.Context) {
	// 1.获取参数及参数校验
	post := new(models.Post)
	if err := context.ShouldBindJSON(post); err != nil {
		zap.L().Debug("context.ShouldBindJSON(post) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(context, CodeInvalidParam)
		return
	}

	// 2.创建帖子
	// 从context中获取到当前登录用户的ID
	userId, err := GetCurrentUser(context)
	if err != nil {
		ResponseError(context, CodeNeedLogin)
		return
	}
	post.AuthorID = userId
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(context, CodeSuccess)
}
