package mux

import (
	"os"

	"github.com/jsjutzi/go-kube-service/api/services/api/mid"
	"github.com/jsjutzi/go-kube-service/api/services/sales/route/sys/checkapi"
	"github.com/jsjutzi/go-kube-service/foundation/logger"
	"github.com/jsjutzi/go-kube-service/foundation/web"
)

func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log))
	checkapi.Routes(mux)

	return mux
}
