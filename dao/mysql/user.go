package mysql

import (
	"BlueBell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret = "payne.com"

func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行sql语句入库
	sqlStr := `insert into user (user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return

}

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// encryptPassword 对密码进行加密
func encryptPassword(oPassword string) string {
	hash := md5.New()
	hash.Write([]byte(secret))

	return hex.EncodeToString(hash.Sum([]byte(oPassword)))
}
