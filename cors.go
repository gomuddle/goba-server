package goba_server

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// globalOptionsHandler returns a handler that
// responds to automatic OPTIONS requests.
func globalOptionsHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "*")
	}
}

// newRouter a new router.Router with a default handler
// for global OPTIONS requests. See globalOptionsHandler.
func newRouter() *router.Router {
	r := router.New()
	r.HandleOPTIONS = true
	r.GlobalOPTIONS = globalOptionsHandler()
	return r
}
