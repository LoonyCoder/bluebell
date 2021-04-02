package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser 获取当前登录用户的userId
func GetCurrentUser(context *gin.Context) (userID int64, err error) {
	uid, ok := context.Get(CtxUserIDKey)

	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
