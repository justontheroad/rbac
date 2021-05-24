package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/justontheroad/rbac/kernel"
)

// Server is a HTTP server wrapper.
type Server struct {
	*http.Server
	lis     net.Listener
	network string
	address string
	timeout time.Duration
	router  *mux.Router
}

func NewServer(configs ...ServerConfig) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: time.Second,
	}
	for _, config := range configs {
		config(srv)
	}
	srv.router = mux.NewRouter()
	srv.Server = &http.Server{Handler: srv}
	return srv
}

func (srv *Server) Handle(path string, h http.Handler) *mux.Route {
	return srv.router.Handle(path, h)
}

func (srv *Server) Subrouters(path string) *mux.Router {
	return srv.router.PathPrefix(path).Subrouter()
}

func (srv *Server) HandlePrefix(path string, h http.Handler) *mux.Route {
	return srv.router.PathPrefix(path).Handler(h)
}

func (srv *Server) HandleFunc(path string, h http.HandlerFunc) *mux.Route {
	return srv.router.HandleFunc(path, h)
}

func (srv *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), srv.timeout)
	defer cancel()
	ctx = kernel.NewContext(ctx, kernel.Transport{Kind: kernel.Http})
	srv.router.ServeHTTP(res, req.WithContext(ctx))
}

func (srv *Server) Start() error {
	lis, err := net.Listen(srv.network, srv.address)
	if err != nil {
		return err
	}
	srv.lis = lis

	if err := srv.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (srv Server) Stop() error {
	return srv.Shutdown(context.Background())
}
