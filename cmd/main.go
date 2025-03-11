package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/pei223/hook-scheduler/internal/domain/hook"
	"github.com/pei223/hook-scheduler/internal/domain/hookschedule"
	"github.com/pei223/hook-scheduler/internal/usecase"
	"github.com/pei223/hook-scheduler/internal/webapi"
	"github.com/pei223/hook-scheduler/internal/worker"
	"github.com/pei223/hook-scheduler/pkg/common"
	"resty.dev/v3"
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

	apiClient := resty.New()

	hookSvc := hook.NewHookService(db, hook.NewHookRepo(), apiClient)
	hookScheduleSvc := hookschedule.NewHookScheduleService(db, hookschedule.NewHookScheduleRepo())

	hookUsecase := usecase.NewHookUsecase(hookSvc)
	hookExecUsecase := usecase.NewHookExecUsecase(db, hookSvc)
	hookScheduleUsecase := usecase.NewHookScheduleUsecase(db, hookScheduleSvc)

	hookRouter := webapi.NewHookRouter(hookUsecase)
	hookScheduleRouter := webapi.NewHookScheduleRouter(hookScheduleUsecase)

	router := webapi.NewRouter(
		hookRouter,
		hookScheduleRouter,
		logger,
	)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ApiServerPort),
		Handler: router,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	invoker := worker.NewInvoker(hookExecUsecase)

	go func() {
		defer stop()

		logger.Info().Msg("listen server")
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				logger.Error().Err(err).Msg("failed to start server")
				syscall.Exit(1)
				return
			}
		}
	}()

	go func() {
		err := invoker.Start(ctx)
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				logger.Error().Err(err).Msg("failed to start invoker")
				syscall.Exit(1)
			}
		}
		stop()
	}()

	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("failed to shutdown server")
		syscall.Exit(1)
		return
	}
	logger.Info().Msg("shutdown succeeded")
}
