package worker

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/pei223/hook-scheduler/internal/hook"
// 	"github.com/rs/zerolog"
// )

// type Invoker struct {
// 	hookMod hook.HookMod
// 	logger  *zerolog.Logger
// }

// func NewInvoker(hookMod hook.HookMod, logger *zerolog.Logger) *Invoker {
// 	return &Invoker{
// 		hookMod: hookMod,
// 		logger:  logger,
// 	}
// }

// func (i *Invoker) Start(ctx context.Context) {
// 	ticker := time.NewTicker(1 * time.Minute)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			i.executeHooks(ctx)
// 		case <-ctx.Done():
// 			log.Println("Invoker stopped")
// 			return
// 		}
// 	}
// }

// func (i *Invoker) executeHooks(ctx context.Context) {
// 	// Fetch all hooks (this is a placeholder logic, replace with actual fetching logic)
// 	hookIds := []uuid.UUID{uuid.New(), uuid.New()} // Replace with actual logic to fetch all hook IDs

// 	for _, hookId := range hookIds {
// 		hook, err := i.hookMod.GetHook(ctx, hookId)
// 		if err != nil {
// 			i.logger.Error().Err(err).Msgf("Failed to get hook: %v", hookId)
// 			continue
// 		}

// 		// Execute the hook (placeholder logic)
// 		i.logger.Info().Msgf("Executing hook: %v", hook)
// 	}
// }
