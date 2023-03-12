package webserver

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/pkg/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hhandhuan/ku-bbs/internal/route"
	"github.com/hhandhuan/ku-bbs/pkg/utils"
)

func Run() {
	gin.SetMode(config.Conf.System.Env)

	engine := gin.Default()
	engine.SetFuncMap(utils.GetTemplateFuncMap())
	engine.Static("/assets", "../assets")
	engine.LoadHTMLGlob("../views/**/**/*")

	store := cookie.NewStore([]byte(config.Conf.Session.Secret))
	engine.Use(sessions.Sessions(config.Conf.Session.Name, store))

	route.RegisterBackendRoute(engine)
	route.RegisterFrontedRoute(engine)

	server := http.Server{Addr: config.Conf.System.Addr, Handler: engine}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("server run error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Conf.System.ShutdownWaitTime)*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("stop server error: ", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("server exiting")
}
