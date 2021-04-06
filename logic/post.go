package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {

	// 1.生成postId

	post.ID = snowflake.GenID()

	// 2.保存到数据库

	return mysql.CreatePost(post)

	// 3.返回

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

func GetPostList() (postList []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList()
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
