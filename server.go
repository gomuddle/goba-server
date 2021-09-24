package goba_server

import (
	"github.com/fasthttp/router"
	"github.com/gomuddle/goba"
	"github.com/valyala/fasthttp"
)

// A Server provides a web interface for interacting with a goba.Goba.
type Server struct {
	goba        *goba.Goba
	router      *router.Router
	credentials []Credentials
}

// New creates a new configured Server and returns it.
func New(goba goba.Goba, validCredentials ...Credentials) *Server {
	s := Server{
		goba:        &goba,
		router:      router.New(),
		credentials: validCredentials,
	}
	s.routes()
	return &s
}

// ListenAndServe  serves HTTP requests from the given TCP address.
func (s Server) ListenAndServe(addr string) error {
	return fasthttp.ListenAndServe(addr, s.router.Handler)
}

// ListenAndServeTLS serves HTTP requests from the given TCP address.
// certFile and keyFile are paths to TLS certificate and key files.
func (s Server) ListenAndServeTLS(addr, certFile, keyFile string) error {
	return fasthttp.ListenAndServeTLS(addr, certFile, keyFile, s.router.Handler)
}
