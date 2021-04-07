package redis

import "BlueBell/models"

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 1.根据用户请求中鞋带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	// 2.确定查询的索引起始位
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 3.ZRevRange查询
	return client.ZRevRange(key, start, end).Result()

}
