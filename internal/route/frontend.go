package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/app/frontend"
)

func RegisterFrontedRoute(engine *gin.Engine) {
	group := engine.Group("/")
	// 社区首页
	group.GET("/", frontend.Home.HomePage)

	// 用户注册
	group.GET("/register", frontend.Auth.RegisterPage)
	// 提交注册
	group.POST("/register", frontend.Auth.RegisterSubmit)

	// 用户登录
	group.GET("/login", frontend.Auth.LoginPage)
	// 登录提交
	group.POST("/login", frontend.Auth.LoginSubmit)
	// 登出用户
	group.GET("/logout", frontend.Auth.LogoutSubmit)

	// 话题发布
	group.GET("/publish", frontend.Topic.PublishPage)
	// 话题提交
	group.POST("/publish", frontend.Topic.PublishSubmit)
	// 话题详情
	group.GET("/topics/:id", frontend.Topic.DetailPage)

	// 评论话题
	group.POST("/comments", frontend.Comment.PublishSubmit)

	// 用户中心
	group.GET("/user", frontend.User.HomePage)
	// 用户编辑
	group.GET("/user/edit", frontend.User.EditPage)
	// 编辑提交
	group.POST("/user/edit", frontend.User.EditSubmit)

	group.POST("/md-upload", frontend.File.MDUploadSubmit)

	// 用户通知
	group.GET("/notice", frontend.Notice.HomePage)

	// 用户点赞
	group.POST("/likes", frontend.Like.LikeSubmit)
	// 用户关注
	group.POST("/follows", frontend.Follow.FollowSubmit)
	// 用户签到
	group.POST("/checkins", frontend.Checkin.StoreSubmit)
}
