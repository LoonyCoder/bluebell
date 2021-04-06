package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的community并返回
	return mysql.GetCommunityList()
}

func GetCommunityListDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityListById(id)
}
