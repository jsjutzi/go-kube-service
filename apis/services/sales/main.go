package main

import (
	"context"
	"os"

	"github.com/jsjutzi/go-kube-service/foundation/logger"
)

func main() {
	// ---------- Logger Setup -------------//
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT ********")
		},
	}

	traceIDFn := func(ctx context.Context) string {
		return "" // web.GetTraceID(ctx)
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES", traceIDFn, events)

	// --------- Logger Setup -------------//

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	return nil
}
