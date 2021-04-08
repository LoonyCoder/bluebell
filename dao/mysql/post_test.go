package mysql

import (
	"BlueBell/models"
	"BlueBell/settings"
	"testing"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "1.15.97.250",
		User:         "root",
		Password:     "root",
		DB:           "bluebell",
		Port:         3306,
		MaxOpenConns: 0,
		MaxIdleConns: 0,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insetr record into mysql failed , err : %v\n", err)
	}
	t.Logf("CreatePost insetr record into mysql success")
}
