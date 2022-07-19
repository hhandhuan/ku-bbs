package backend

import "github.com/hhandhuan/ku-bbs/internal/model"

type GetUserListReq struct {
	Page     int    `form:"page"`
	Keywords string `form:"keywords"`
}

type User struct {
	model.Users
}
