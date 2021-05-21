package kernal

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
	lis     net.Listener
	network string
	address string
	timeout time.Duration
	router  *mux.Router
}

func NewServer() *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: time.Second,
	}
	srv.router = mux.NewRouter()
	srv.Server = &http.Server{Handler: srv}
	return srv
}

func (srv *Server) Hnadler(path string, h http.Handler) {
	srv.router.Handle(path, h)
}

func (srv *Server) HandlePrefix(path string, h http.Handler) {
	srv.router.PathPrefix(path).Handler(h)
}

func (srv *Server) HnadlerFun(path string, h http.HandlerFunc) {
	srv.router.HandleFunc(path, h)
}

func (srv *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), srv.timeout)
	defer cancel()
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
