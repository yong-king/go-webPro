package logic

import (
	"bubble/dao/mysql"
	"bubble/models"
)

func CommunityList() (communityList []*models.Community, err error) {
	return mysql.CommunityList()
}

func CommunityDetailData(id int64) (communityDetailData *models.CommunityDetail, err error) {
	return mysql.CommunityDetailDataById(id)
}
