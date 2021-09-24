package goba_server

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

// handlerFunc is the type of any endpoint handler func.
type handlerFunc func(ctx *fasthttp.RequestCtx) error

// handle wraps the given handler.
func (s Server) handle(handler handlerFunc) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		globalOptionsHandler()(ctx)
		if err := handler(ctx); err != nil {
			s.error(ctx, err, http.StatusInternalServerError)
		}
	}
}
