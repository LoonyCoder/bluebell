package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询异常
		return err
	}
	// 生成UID
	userId := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) error {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
