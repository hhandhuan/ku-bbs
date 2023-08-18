package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hhandhuan/ku-bbs/pkg/logger"
	"github.com/hhandhuan/ku-bbs/pkg/mysql"
	"github.com/hhandhuan/ku-bbs/pkg/redis"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/ku-bbs/pkg/config"

	"github.com/hhandhuan/ku-bbs/internal/route"
	"github.com/hhandhuan/ku-bbs/pkg/utils"
)

func init() {
	config.Initialize()

	logger.Initialize(config.GetInstance().Logger)
	mysql.Initialize(config.GetInstance().Mysql)
	redis.Initialize(config.GetInstance().Redis)

	gin.SetMode(config.GetInstance().System.Env)
}

func main() {
	logger.GetInstance().Info().Msg("service is starting")

	engine := gin.Default()
	engine.SetFuncMap(utils.GlobalFunc())
	engine.Static("/assets", "../assets")
	engine.LoadHTMLGlob("../views/**/**/*")

	store := cookie.NewStore([]byte(config.GetInstance().Session.Secret))
	engine.Use(sessions.Sessions(config.GetInstance().Session.Name, store))

	route.RegisterBackendRoute(engine)
	route.RegisterFrontedRoute(engine)

	server := http.Server{Addr: config.GetInstance().System.Addr, Handler: engine}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.GetInstance().Error().Msgf("listen server error: %v", err)
		}
	}()

	logger.GetInstance().Info().Msg("service started successfully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.GetInstance().System.ShutdownWaitTime)*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.GetInstance().Error().Msgf("service shutdown error: %v", err)
	}

	<-ctx.Done()

	logger.GetInstance().Info().Msg("service has been shut down")
}
