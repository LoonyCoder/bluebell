package logic

import (
	"BlueBell/dao/redis"
	"BlueBell/models"
	"go.uber.org/zap"
	"strconv"
)

// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm

// 投票功能：
// 本项目使用简化版的投票分数算法

/**
	投票的几种情况：
	direction = 1时，有两种情况：
		1. 之前没有投过票，现在投赞成票  	--> 更新分数和投票记录
		2. 之前投反对票，现在改投赞成票  	--> 更新分数和投票记录
	direction = 0时，有两种情况：
		1. 之前投过赞成票，现在要取消投票		--> 更新分数和投票记录
		2. 之前投过反对票，现在要取消投票		--> 更新分数和投票记录
	direction = -1时，有两种情况：
		1. 之前没有投过票，现在投反对票		--> 更新分数和投票记录
		2. 之前投赞成票，现在改投反对票		--> 更新分数和投票记录

投票的限制：
	每个帖子自发表之日起，一个星期之内允许用户投票，超过一个星期就不允许投票了
	到期之后，将redis中保存的赞成票数及反对票数存储到mysql表中
	到期之后，删除 KeyPostVotedZSetPrefix
*/
// VoteForPost 为帖子投票的函数
func VoteForPost(userId int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost", zap.Int64("userId", userId), zap.String("postId", p.PostID), zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userId)), p.PostID, float64(p.Direction))

}
