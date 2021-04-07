package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

/**
	投票的几种情况：
	direction = 1时，有两种情况：
		1. 之前没有投过票，现在投赞成票  	--> 更新分数和投票记录	差值的绝对值：1  +432
		2. 之前投反对票，现在改投赞成票  	--> 更新分数和投票记录	差值的绝对值：2  +432 * 2
	direction = 0时，有两种情况：
		1. 之前投过赞成票，现在要取消投票		--> 更新分数和投票记录	差值的绝对值：1  -432
		2. 之前投过反对票，现在要取消投票		--> 更新分数和投票记录	差值的绝对值：1  +432
	direction = -1时，有两种情况：
		1. 之前没有投过票，现在投反对票		--> 更新分数和投票记录	差值的绝对值：1  -432
		2. 之前投赞成票，现在改投反对票		--> 更新分数和投票记录	差值的绝对值：2  -432 * 2

投票的限制：
	每个帖子自发表之日起，一个星期之内允许用户投票，超过一个星期就不允许投票了
	到期之后，将redis中保存的赞成票数及反对票数存储到mysql表中
	到期之后，删除 KeyPostVotedZSetPrefix
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func VoteForPost(userId, postId string, value float64) error {

	// 1.判断投票的限制 (从redis获取帖子的发布时间)
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postId).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2.更新帖子的分数
	// 先查询当前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postId), userId).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	_, err := client.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postId).Result()
	if ErrVoteTimeExpire != nil {
		return err
	}
	// 3.记录用户为该帖子投过票
	if value == 0 {
		_, err = client.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postId), userId).Result()
	} else {
		_, err = client.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postId), redis.Z{
			Score:  value,
			Member: userId,
		}).Result()
	}
	return err
}