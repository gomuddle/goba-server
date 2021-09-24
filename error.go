package goba_server

import (
	"github.com/valyala/fasthttp"
	"log"
)

// ErrorResponse represents a response to a failed request.
type ErrorResponse struct {
	// Error describes the reason why the request failed.
	Error string `json:"error,omitempty"`
}

// error sets the response's status code to the given status
// code and sends an ErrorResponse with err to the client.
func (s Server) error(ctx *fasthttp.RequestCtx, err error, statuscode int) {
	ctx.SetStatusCode(statuscode)
	resp := ErrorResponse{Error: err.Error()}
	if err = s.writeJSON(ctx, resp); err != nil {
		log.Println("error:", err)
	}
}
