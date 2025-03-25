package models

type Community struct {
	ID   int64  `json:"id,string" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64  `json:"id,string" db:"community_id"`
	Name         string `json:"name" db:"community_name"`
	Introduction string `json:"introduction" db:"introduction"`
	Create_time  string `json:"create_time" db:"create_time"`
	Update_time  string `json:"update_time" db:"update_time"`
}
