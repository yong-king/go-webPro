package controller

import "bubble/models"

type _ResponsePostList struct {
	Code ResCode      `json:"code"` // 业务响应状态码
	Msg  string       `json:"msg"`  // 提示信息，与 Response 结构体对齐
	Data *models.Post `json:"data"` // 数据
}

type _ResponsePostLists struct {
	Code ResCode             `json:"code"` // 业务响应状态码
	Msg  string              `json:"msg"`  // 提示信息，与 Response 结构体对齐
	Data []models.PostDetail `json:"data"` // 帖子详情列表
}

type _ResponseCommunityLists struct {
	Code ResCode            `json:"code"` // 业务响应状态码
	Msg  string             `json:"msg"`  // 提示信息，与 Response 结构体对齐
	Data []models.Community `json:"data"` // 社区列表
}

type _ResponseCommunityDetailLists struct {
	Code ResCode                  `json:"code"` // 业务响应状态码
	Msg  string                   `json:"msg"`  // 提示信息，与 Response 结构体对齐
	Data []models.CommunityDetail `json:"data"` // 社区详细信息
}

type _ResponseUser struct {
	Code ResCode             `json:"code"` // 业务响应状态码
	Msg  string              `json:"msg"`  // 提示信息，与 Response 结构体对齐
	Data *models.ParamSignUp `json:"data"` // 社区详细信息
}

type _ResponseAuth struct {
	Code ResCode      `json:"code"` // 业务响应状态码
	Msg  string       `json:"msg"`  // 提示信息，与 Response 结构体对齐
	Data *models.User `json:"data"` // 社区详细信息
}

type _ResponseVote struct {
	Code ResCode               `json:"code"` // 业务响应状态码
	Msg  string                `json:"msg"`  // 提示信息，与 Response 结构体对齐
	Data *models.ParamVoteData `json:"data"` // // 使用指针，防止空结构体导致返回 nil
}
