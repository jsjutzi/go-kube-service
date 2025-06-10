package mid

import (
	"context"
	"fmt"
	"time"

	"github.com/jsjutzi/go-kube-service/foundation/logger"
	"github.com/jsjutzi/go-kube-service/foundation/web"
)

// Logger executes the handler and logs the request details.
func Logger(ctx context.Context, log *logger.Logger, path string, rawQuery string, method string, remoteAddr string, handler Handler) error {
	v := web.GetValues(ctx)

	if rawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, rawQuery)
	}

	log.Info(ctx, "request started", "method", method, "path", path, "remote_addr", remoteAddr)

	err := handler(ctx)

	log.Info(ctx, "request completed", "method", method, "path", path, "remote_addr", remoteAddr, "statusCode", v.StatusCode, "since", time.Since(v.Now).String())

	return err
}
