package main

import (
	"context"
	"github.com/shiwa-sh/usdl/chat/foundation/logger"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	var log *logger.Logger

	traceIDFn := func(ctx context.Context) string {
		return "" // TODO: NEED TRACE IDs
	}

	log = logger.New(os.Stdout, logger.LevelInfo, "CAP", traceIDFn)

	// -------------------------------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "err", err)
		os.Exit(1)
	}
}
func run(ctx context.Context, log *logger.Logger) error {
	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// ----------------------------------------------------------------------------

	log.Info(ctx, "startup", "status", "starting")
	defer log.Info(ctx, "startup", "status", "shouting down")
	shoutdown := make(chan os.Signal, 1)
	signal.Notify(shoutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shoutdown
	return nil
}
