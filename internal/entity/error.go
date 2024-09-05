package entity

import "kratos-use/api/common"

var (
	ErrUserConflict = common.ErrorConflict("用户已存在")
)
