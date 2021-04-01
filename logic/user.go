package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {

	// 判断用户是否存在
	mysql.QueryUserByUsername()
	// 生成UID
	snowflake.GenID()
	// 密码加密
	// 保存数据库
	mysql.InsertUser()
}
