package controller

import (
	"BlueBell/logic"
	"BlueBell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子的处理函数
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

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(context *gin.Context) {
	// 1.获取参数（从URL中获取帖子的ID）
	pidStr := context.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}

	// 2.根据ID获取帖子的数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(context, data)
}

// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(context *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(context)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(context, data)
}

// GetPostListHandler2 升级版忒自列表接口
// 根据前端传来的排序参数（按创建时间、按分数） 动态获取贴子列表
func GetPostListHandler2(context *gin.Context) {
	// 1.获取参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := context.ShouldBindQuery(p); err != nil {
		zap.L().Error("logic.GetPostListHandler2() with invalid params", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}


	// 获取数据
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(context, data)
}

func getPageInfo(context *gin.Context) (int64, int64) {
	pageStr := context.Query("page")
	sizeStr := context.Query("size")

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 0
	}

	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
