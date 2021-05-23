package app

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type App struct {
	config appConfig
	ctx    context.Context
	cancle func()
	// srv    *Server.Server
	// sigs   []os.Signal
}

func New(config ...AppConfig) *App {
	appconf := appConfig{
		id:      "",
		name:    "",
		version: "",
		ctx:     context.Background(),
		sigs:    []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		// servers: []Server.Server{},
	}
	if id, err := uuid.NewUUID(); err == nil {
		appconf.id = id.String()
	}
	for _, conf := range config {
		conf(&appconf)
	}
	ctx, cancle := context.WithCancel(context.Background())
	return &App{
		config: appconf,
		ctx:    ctx,
		cancle: cancle,
	}
}

func (a *App) Run() error {
	g, ctx := errgroup.WithContext(a.ctx)
	for _, srv := range a.config.servers {
		srv := srv
		g.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		g.Go(func() error {
			return srv.Start()
		})
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, a.config.sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-s:
				a.Stop()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (a *App) Stop() error {
	if a.cancle != nil {
		a.cancle()
	}

	return nil
}
