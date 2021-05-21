package rbac

import (
	"context"
	"errors"
	"os"
	"os/signal"

	Server "github.com/justontheroad/rbac/kernal/server"
	"golang.org/x/sync/errgroup"
)

type App struct {
	ctx    context.Context
	cancle func()
	srv    *Server.Server
	sigs   []os.Signal
}

func New() *App {
	ctx, cancle := context.WithCancel(context.Background())
	srv := Server.NewServer()
	return &App{
		ctx:    ctx,
		cancle: cancle,
		srv:    srv,
	}
}

func (a *App) Run() error {
	g, ctx := errgroup.WithContext(a.ctx)
	g.Go(func() error {
		<-ctx.Done()
		return a.srv.Stop()
	})
	g.Go(func() error {
		return a.srv.Start()
	})

	s := make(chan os.Signal, 1)
	signal.Notify(s, a.sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-s:
				a.srv.Stop()
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
