package server

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Server has a http server instance
type Server struct {
	srv *http.Server
}

// New returns a pointer to a populated Server object
func New() *Server {
	return &Server{srv: &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}}
}

// WithAddr returns a Server instance with the Addr initialized
func (s *Server) WithAddr(addr string) *Server {
	s.srv.Addr = addr
	return s
}

// WithErrLogger returns a pointer to a server instance with a new ErrorLog
func (s *Server) WithErrLogger(l *log.Logger) *Server {
	s.srv.ErrorLog = l
	return s
}

// WithRouter returns a pointer to a server instance with a new handler
func (s *Server) WithRouter(router *chi.Mux) *Server {
	s.srv.Handler = router
	return s
}

// Start starts the sever
func (s *Server) Start() error {
	if len(s.srv.Addr) == 0 {
		return errors.New("server missing address")
	}

	if s.srv.Handler == nil {
		return errors.New("server missing handler")
	}

	return s.srv.ListenAndServe()
}

// Close closes the server connection
func (s *Server) Close() error {
	return s.srv.Close()
}
