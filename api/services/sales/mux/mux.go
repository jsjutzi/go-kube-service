package mux

import (
	"os"

	"github.com/jsjutzi/go-kube-service/api/services/sales/route/sys/checkapi"
	"github.com/jsjutzi/go-kube-service/foundation/web"
)

func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)
	checkapi.Routes(mux)

	return mux
}
