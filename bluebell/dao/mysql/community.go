package mysql

import (
	"bubble/models"
	"database/sql"
	"go.uber.org/zap"
)

func CommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("社区数据库为空")
			err = nil
		}
	}
	return
}

func CommunityDetailDataById(id int64) (communityDetailData *models.CommunityDetail, err error) {
	sqlStr := "select community_id, community_name, introduction, create_time, update_time from community where community_id = ?"
	communityDetailData = new(models.CommunityDetail)
	err = db.Get(communityDetailData, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
		zap.L().Error("db.Select(communityDetailData, sqlStr, id) failed", zap.Error(err))
	}
	return
}

// GetCommunityNameById 根据id获取社区名称
func GetCommunityNameById(id int64) (community *models.CommunityDetail, err error) {
	sqlStr := "select community_name from community where community_id = ?"
	community = new(models.CommunityDetail)
	err = db.Get(community, sqlStr, id)
	return
}
