package gobaserver

import (
	"errors"
	"github.com/valyala/fasthttp"
)

// ErrInvalidCredentials occurs when credentials have been
// provided to an endpoint that are not known to the server.
var ErrInvalidCredentials = errors.New("invalid credentials")

// authMiddleware validates the request's authorization
// header and calls the given handler if they are valid.
func (s Server) authMiddleware(handler handlerFunc) handlerFunc {
	return func(ctx *fasthttp.RequestCtx) error {
		authHeader := ctx.Request.Header.Peek("Authorization")
		creds, err := credentialsFromHeader(string(authHeader))
		if err != nil {
			return err
		}
		if !s.credentialsValid(*creds) {
			return ErrInvalidCredentials
		}
		return handler(ctx)
	}
}

// credentialsValid reports whether the given credentials are known to the server.
func (s Server) credentialsValid(creds Credentials) bool {
	for _, c := range s.credentials {
		if c.Username == creds.Username && c.Password == creds.Password {
			return true
		}
	}
	return false
}
