package checkapi

import "github.com/jsjutzi/go-kube-service/foundation/web"

func Routes(mux *web.App) {
	mux.HandleFunc("GET /v1/liveness", liveness)
	mux.HandleFunc("GET /v1/readiness", readiness)
	mux.HandleFunc("GET /v1/testerror", testError)
}
