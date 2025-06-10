package checkapi

import "github.com/jsjutzi/go-kube-service/foundation/web"

func Routes(mux *web.App) {
	mux.HandleFuncWithNoMiddleware("GET /v1/liveness", liveness)
	mux.HandleFuncWithNoMiddleware("GET /v1/readiness", readiness)
	mux.HandleFunc("GET /v1/testerror", testError)
	mux.HandleFunc("GET /v1/testpanic", testPanic)
}
