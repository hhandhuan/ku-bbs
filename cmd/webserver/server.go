package webserver

import (
	"github.com/huhaophp/hblog/pkg/config"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/huhaophp/hblog/internal/route"
	_ "github.com/huhaophp/hblog/pkg/config"
	_ "github.com/huhaophp/hblog/pkg/db"
	_ "github.com/huhaophp/hblog/pkg/redis"
	"github.com/huhaophp/hblog/pkg/utils"
)

func Run() {
	engine := gin.Default()

	engine.SetFuncMap(utils.GetTemplateFuncMap())

	engine.Static("/a", "../assets/a")
	engine.Static("/u", "../assets/u")
	engine.LoadHTMLGlob("../views/**/**/*")

	store := cookie.NewStore([]byte(config.Conf.Session.Secret))
	engine.Use(sessions.Sessions(config.Conf.Session.Name, store))

	route.RegisterBackendRoute(engine)
	route.RegisterFrontedRoute(engine)

	if err := engine.Run(":8081"); err != nil {
		log.Fatalf("server running error: %v", err)
	}
}
