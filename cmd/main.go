package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/pei223/hook-scheduler/internal/task"
	"github.com/pei223/hook-scheduler/internal/web"
	"github.com/pei223/hook-scheduler/pkg/common"
)

type Config struct {
	ApiServerPort            string `envconfig:"API_SERVER_PORT" default:"80"`
	DatabaseConnectionString string `envconfig:"DATABASE_CONNECTION_STRING" default:""`
	LogLevel                 string `envconfig:"LOG_LEVEL" default:"info"`
}

func main() {
	Serve()
}

func Serve() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	logger := common.NewLogger(ctx, cfg.LogLevel)
	db, err := sql.Open("postgres", cfg.DatabaseConnectionString)
	if err != nil {
		panic(err)
	}

	taskMod := task.NewTaskMod(task.NewTaskRepo(db))
	taskWebHdlr := task.NewTaskWebHandler(&logger, taskMod)

	router := web.NewRouter(
		taskWebHdlr,
	)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ApiServerPort),
		Handler: router,
	}

	go func() {
		defer stop()

		logger.Info().Msg("listen server")
		if err := server.ListenAndServe(); err != nil {
			logger.Error().Err(err).Msg("failed to start server")
			syscall.Exit(1)
			return
		}
	}()
	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("failed to shutdown server")
		syscall.Exit(1)
		return
	}
	logger.Info().Msg("shutdown succeeded")
}
