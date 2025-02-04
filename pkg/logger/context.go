package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type loggerKey struct{}

func NewLogger(ctx context.Context,
	logLevel string,
) zerolog.Logger {
	return zerolog.New(os.Stdout).
		Level(lo.Must(zerolog.ParseLevel(logLevel))).With().
		Ctx(ctx).
		Timestamp().
		Caller().
		Logger()
}

func FromContext(ctx context.Context) zerolog.Logger {
	if ctx == nil {
		logger := NewLogger(ctx, "info")
		logger.Warn().Msg("no logger from context")
		return logger
	}
	if logger, ok := ctx.Value(loggerKey{}).(zerolog.Logger); ok {
		return logger
	}
	logger := NewLogger(ctx, "info")
	logger.Warn().Msg("no logger from context")
	return logger
}

func WithContext(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}
