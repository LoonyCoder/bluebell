package mysql

import "BlueBell/models"

func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post (
				post_id,title,content,author_id,community_id)
				values(?,?,?,?,?)`

	_, err = db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return
}

func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id , title,content , author_id , community_id , create_time
				from post
				where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}
