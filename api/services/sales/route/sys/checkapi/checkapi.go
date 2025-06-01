package checkapi

import (
	"context"
	"net/http"

	"github.com/jsjutzi/go-kube-service/foundation/web"
)

func liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "ok",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "ready",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
