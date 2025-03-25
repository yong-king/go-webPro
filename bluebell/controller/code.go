package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeInvalidPassword
	CodeUserNotExit
	CodeUserExit
	CodeServeBusy

	CodeNeedLogin
	CodeInvalidToken

	CodeMutilUser

	CodeServerBusy
)

var codeMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "参数错误",
	CodeInvalidPassword: "用户名或密码错误",
	CodeUserNotExit:     "用户不存在",
	CodeUserExit:        "用户已存在",
	CodeServeBusy:       "服务繁忙",

	CodeNeedLogin:    "需要用户登录",
	CodeInvalidToken: "非法token",

	CodeMutilUser: "多个用户登录",

	CodeServerBusy: "社区繁忙",
}

// GetMsg 错误信息
func (c ResCode) GetMsg() string {
	msg, ok := codeMap[c]
	if !ok {
		msg = codeMap[CodeInvalidParam]
	}
	return msg
}
