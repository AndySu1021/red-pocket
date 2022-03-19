package cmd

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"red-packet/config"
	"red-packet/pkg/cache"
	"red-packet/pkg/repository"
	"red-packet/pkg/restful"
	"red-packet/pkg/service"
	"red-packet/util/db"
	"red-packet/util/gin"
	"red-packet/util/helper"
	"red-packet/util/lock"
	"red-packet/util/redis"
	zlog "red-packet/util/zerolog"
	"syscall"
	"time"
)

// ServerCmd 是此程式的Service入口點
var ServerCmd = &cobra.Command{
	Run: run,
	Use: "server",
}

var commonModule = fx.Options(
	fx.Provide(
		config.NewConfig,
		gin.NewGin,
		db.NewDatabase,
		redis.NewRedisClient,
		lock.NewRedisLocker,
	),
	fx.Invoke(
		zlog.Init,
		restful.RegisterAPIRouter,
	),
)

var otherModule = fx.Options(
	repository.Module,
	cache.Module,
	restful.Module,
	service.Module,
)

func run(_ *cobra.Command, _ []string) {
	defer helper.Recover(context.Background())

	logger := log.Level(zerolog.InfoLevel)
	fxOption := []fx.Option{
		fx.Logger(&logger),
	}

	fxOption = append(fxOption, commonModule, otherModule)

	app := fx.New(
		fxOption...,
	)

	exitCode := 0
	if err := app.Start(context.Background()); err != nil {
		log.Err(err).Msg("app start err")
		os.Exit(exitCode)
		return
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-stopChan
	log.Info().Msgf("main: shutting down server...")

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Err(err).Msg("app stop err")
	}

	os.Exit(exitCode)
}
