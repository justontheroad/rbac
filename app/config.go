package app

import (
	"context"
	"os"

	kernel "github.com/justontheroad/rbac/kernel"
)

type AppConfig func(*appConfig)

type appConfig struct {
	id      string
	name    string
	version string
	ctx     context.Context
	sigs    []os.Signal
	servers []kernel.Server
}

func SetID(id string) AppConfig {
	return func(config *appConfig) {
		config.id = id
	}
}

func SetName(name string) AppConfig {
	return func(config *appConfig) {
		config.name = name
	}
}

func SetVersion(version string) AppConfig {
	return func(config *appConfig) {
		config.version = version
	}
}

func SetContext(ctx context.Context) AppConfig {
	return func(config *appConfig) {
		config.ctx = ctx
	}
}

func SetSignal(sigs ...os.Signal) AppConfig {
	return func(config *appConfig) {
		config.sigs = sigs
	}
}

func SetServer(servers ...kernel.Server) AppConfig {
	return func(config *appConfig) {
		config.servers = servers
	}
}
