package common

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

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
