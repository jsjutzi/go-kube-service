package mid

import (
	"context"
	"fmt"

	"github.com/jsjutzi/go-kube-service/foundation/logger"
)

func Logger(ctx context.Context, log *logger.Logger, path string, rawQuery string, method string, remoteAddr string, handler Handler) error {
	if rawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, rawQuery)
	}
	log.Info(ctx, "request started", "method", method, "path", path, "remote_addr", remoteAddr)

	err := handler(ctx)

	log.Info(ctx, "request completed", "method", method, "path", path, "remote_addr", remoteAddr)

	return err
}
