package frontend

import "github.com/huhaophp/hblog/internal/model"

type LoginReq struct {
	Name     string `v:"required|regex:[\u4e00-\u9fa5a-zA-Z0-9]{2,8}#用户名错误|用户名格式错误" form:"name"`
	Password string `v:"required|length:6,20#账号或密码错误|密码长度错误" form:"password"`
}

type RegisterReq struct {
	Name            string `v:"required|regex:[\u4e00-\u9fa5a-zA-Z0-9]{2,8}#用户名错误|用户名格式错误" form:"name"`
	Password        string `v:"required|length:6,20|same:confirm_password#账号或密码错误|密码长度错误|密码和确认密码不一致" form:"password"`
	ConfirmPassword string `v:"required|length:6,20#确认密码错误|密码长度错误" form:"confirm_password"`
}

type EditUserReq struct {
	Name   string `v:"required|regex:[\u4e00-\u9fa5a-zA-Z0-9]{2,8}#用户名错误|用户名格式错误" form:"name"`
	Gender int    `form:"gender"`
	Email  string `form:"email"`
	City   string `form:"city"`
	Site   string `form:"site"`
	Job    string `form:"job"`
	Desc   string `form:"desc"`
	Tab    string `form:"tab"`
}

type EditPasswordReq struct {
	OldPassword     string `v:"required|length:6,20#旧密码错误|旧密码错误" form:"old_password"`
	Password        string `v:"required|length:6,20|same:confirm_password#密码错误|密码长度错误|密码和确认密码不一致" form:"password"`
	ConfirmPassword string `v:"required|length:6,20#确认密码错误|密码长度错误" form:"confirm_password"`
}

type GetUserHomeReq struct {
	Page int    `form:"page"`
	Tab  string `form:"tab"`
	ID   int    `form:"id"`
}

type FollowUserReq struct {
	UserID uint64 `v:"required#关注目标错误" form:"user_id"`
}

type User struct {
	model.Users
	Follow  *model.Follows  `gorm:"foreignKey:target_id"`
	Checkin *model.Checkins `gorm:"foreignKey:user_id"`
}
