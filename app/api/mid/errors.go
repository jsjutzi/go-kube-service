package mid

import (
	"context"

	"github.com/jsjutzi/go-kube-service/app/api/errs"
	"github.com/jsjutzi/go-kube-service/foundation/logger"
)

func Errors(ctx context.Context, log *logger.Logger, handler Handler) error {
	err := handler(ctx)

	if err == nil {
		return nil
	}

	log.Error(ctx, "message", "error", err.Error())

	if errs.IsError(err) {
		return errs.GetError(err)
	}

	return errs.Newf(errs.Unknown, errs.Unknown.String())
}
