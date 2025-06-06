package mid

import (
	"context"
	"net/http"

	"github.com/jsjutzi/go-kube-service/app/api/mid"
	"github.com/jsjutzi/go-kube-service/foundation/logger"
	"github.com/jsjutzi/go-kube-service/foundation/web"
)

func Logger(log *logger.Logger) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Logger(ctx, log, r.URL.Path, r.URL.RawQuery, r.Method, r.RemoteAddr, hdl)
		}

		return h
	}

	return m
}
