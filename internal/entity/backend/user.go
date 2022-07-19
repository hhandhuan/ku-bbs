package backend

import "github.com/huhaophp/hblog/internal/model"

type GetUserListReq struct {
	Page     int    `form:"page"`
	Keywords string `form:"keywords"`
}

type User struct {
	model.Users
}
