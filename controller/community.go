package controller

import (
	"BlueBell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//  社区相关

func CommunityHandler(context *gin.Context) {

	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(context, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(context *gin.Context) {
	// 获取社区id
	communityId := context.Param("id")
	id, err := strconv.ParseInt(communityId, 10, 64)
	if err != nil {
		ResponseError(context,CodeInvalidParam)
	}
	// 查询到所有的社区（community_id,community_name）以列表的形式返回
	data, err := logic.GetCommunityListDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityListDetail() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(context, data)
}
