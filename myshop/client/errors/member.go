package errors

import "errors"

var (
	UserExists = errors.New("用户已存在")
	CreateUserFail = errors.New("创建用户失败")
	UserUnExist = errors.New("用户不存在")
	TokenParseError = errors.New("token 解析 错误")
	ObjectCheckError = errors.New("断言失败")
)
