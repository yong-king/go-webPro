package mysql

import (
	"bubble/models"
	"bubble/setting"
	"testing"
)

func init() {
	config := &setting.MysqlConfig{
		Host:     "127.0.0.1",
		User:     "root",
		Password: "youngking98",
		DB:       "bluebell",
		Port:     3307,
		MaxIdle:  10,
		MaxOpen:  10,
	}
	err := Init(config)
	if err != nil {
		panic(err)
	}
}

func TestPostData(t *testing.T) {
	post := &models.Post{
		ID:          5,
		AuthorID:    11,
		CommunityId: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := PostData(post)
	if err != nil {
		t.Fatalf("PostData() failed, err: %v", err)
	}
	t.Logf("PostData() success")
}
