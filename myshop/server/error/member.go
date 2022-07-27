package error

import "errors"

var (
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorSearch = errors.New("查询异常")
)
