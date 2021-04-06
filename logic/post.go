package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
)

func CreatePost(post *models.Post) (err error) {

	// 1.生成postId

	post.ID = snowflake.GenID()

	// 2.保存到数据库

	return mysql.CreatePost(post)

	// 3.返回

}
