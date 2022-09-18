package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/internal/app/backend"
)

func RegisterBackendRoute(engine *gin.Engine) {
	group := engine.Group("backend")

	group.Use(isAdmin)

	group.GET("/", backend.Home.IndexPage)
	group.GET("/users", backend.User.IndexPage)

	group.GET("/topics", backend.Topic.IndexPage)
	group.POST("/topics/:id", backend.Topic.DeleteSubmit)

	group.GET("notices", backend.Notice.IndexPage)
	group.GET("notices/publish", backend.Notice.PublishPage)
	group.POST("notices/publish", backend.Notice.PublishSubmit)

	group.GET("nodes", backend.Node.IndexPage)
	group.GET("nodes/create", backend.Node.CreatePage)
	group.POST("nodes/create", backend.Node.CreateSubmit)
	group.GET("nodes/:id/edit", backend.Node.EditPage)
	group.POST("nodes/:id/edit", backend.Node.EditSubmit)

}
