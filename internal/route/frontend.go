package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/app/frontend"
)

func RegisterFrontedRoute(engine *gin.Engine) {
	group := engine.Group("/")

	// 用户注册
	group.GET("/register", frontend.Auth.RegisterPage)
	// 提交注册
	group.POST("/register", frontend.Auth.RegisterSubmit)

	// 用户登录
	group.GET("/login", frontend.Auth.LoginPage)
	// 登录提交
	group.POST("/login", frontend.Auth.LoginSubmit)

	group.Use(visitor)

	// 社区首页
	group.GET("/", frontend.Home.HomePage)

	// 登出用户
	group.GET("/logout", frontend.Auth.LogoutSubmit)

	// 话题发布
	group.GET("/publish", frontend.Topic.PublishPage)
	// 话题提交
	group.POST("/publish", frontend.Topic.PublishSubmit)

	// 话题详情
	group.GET("/topics/:id", frontend.Topic.DetailPage)
	// 删除话题
	group.POST("/topics/:id/delete", frontend.Topic.DeleteSubmit)
	// 编辑话题
	group.GET("/topics/:id/edit", frontend.Topic.EditPage)
	// 编辑提交
	group.POST("/topics/:id/edit", frontend.Topic.EditSubmit)
	// 编辑讨论状态
	group.POST("/topics/:id/comment-state", frontend.Topic.SettingCommentStateSubmit)

	// 评论话题
	group.POST("/comments", frontend.Comment.PublishSubmit)
	// 删除评论
	group.POST("comments/delete", frontend.Comment.DeleteSubmit)

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

	// 举报资源
	group.POST("/reports", frontend.Report.ReportSubmit)

	// 检索列表
	group.GET("/search", frontend.Search.List)
}
