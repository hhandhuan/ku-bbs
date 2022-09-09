package webserver

import (
	"github.com/hhandhuan/ku-bbs/pkg/config"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/hhandhuan/ku-bbs/internal/route"
	"github.com/hhandhuan/ku-bbs/pkg/utils"
)

func Run() {
	engine := gin.Default()

	engine.SetFuncMap(utils.GetTemplateFuncMap())

	engine.Static("/assets", "../assets")
	engine.LoadHTMLGlob("../views/**/**/*")

	store := cookie.NewStore([]byte(config.Conf.Session.Secret))
	engine.Use(sessions.Sessions(config.Conf.Session.Name, store))

	route.RegisterBackendRoute(engine)
	route.RegisterFrontedRoute(engine)

	if err := engine.Run(":8081"); err != nil {
		log.Fatalf("server running error: %v", err)
	}
}
