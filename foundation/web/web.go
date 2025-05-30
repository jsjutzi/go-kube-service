package web

import (
	"context"
	"net/http"
	"os"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
	shutdown chan os.Signal
	mw       []MidHandler
}

// NewApp creates a new web application with a ServeMux and a shutdown channel.
func NewApp(shutdown chan os.Signal, mw ...MidHandler) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
		mw:       mw,
	}
}

// HandleFunc allows us to override the default ServeMux HandleFunc method to use our custom Handler type.
func (a *App) HandleFunc(pattern string, handler Handler, mw ...MidHandler) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		handler(r.Context(), w, r)
	}

	a.ServeMux.HandleFunc(pattern, h)
}
