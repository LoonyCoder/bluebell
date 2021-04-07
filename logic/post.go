package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/dao/redis"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {

	// 1.生成postId

	post.ID = snowflake.GenID()

	// 2.保存到数据库

	err = mysql.CreatePost(post)
	if err != nil {
		return err
	}
	err = redis.CreatePost(post.ID)
	// 3.返回
	return
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {

	// 查询并组合接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 根据作者ID 查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.getUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	// 根据社区ID 查询社区信息
	community, err := mysql.GetCommunityListById(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityListById(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		CommunityDetail: community,
		Post:            post,
	}
	return
}

func GetPostList(page, size int64) (postList []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	postList = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 根据作者ID 查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.getUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		// 根据社区ID 查询社区信息
		community, err := mysql.GetCommunityListById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityListById(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			CommunityDetail: community,
			Post:            post,
		}
		postList = append(postList, postDetail)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (postList []*models.ApiPostDetail, err error) {
	// 2.从redis查询id列表
	ids, err := redis.GetPostIdsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIdsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	// 3.根据id去数据库查询贴子详细信息
	posts, err := mysql.GetPostListByIds(ids)
	if err != nil {
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
	// 将贴子的作者和社区信息查询出来填充到贴子中
	for _, post := range posts {
		// 根据作者ID 查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.getUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		// 根据社区ID 查询社区信息
		community, err := mysql.GetCommunityListById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityListById(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			CommunityDetail: community,
			Post:            post,
		}
		postList = append(postList, postDetail)
	}
	return

}
