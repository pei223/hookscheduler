package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/pei223/hook-scheduler/internal/task"
	"github.com/pei223/hook-scheduler/internal/web"
	"github.com/pei223/hook-scheduler/pkg/common"
)

func main() {
	Serve()
}

func Serve() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	logger := common.NewLogger(ctx, "debug")
	taskWebHdlr := task.NewTaskWebHandler(&logger)
	router := web.NewRouter(
		taskWebHdlr,
	)
	server := &http.Server{
		Addr:    ":80", // TODO 環境変数
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
