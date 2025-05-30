package checkapi

import "github.com/jsjutzi/go-kube-service/foundation/web"

func Routes(mux *web.App) {
	mux.HandleFunc("GET /liveness", liveness)
	mux.HandleFunc("GET /readiness", readiness)
}
