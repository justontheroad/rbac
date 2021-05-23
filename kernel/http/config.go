package http

import (
	"time"
)

type ServerConfig func(*Server)

func SetNetwork(network string) ServerConfig {
	return func(s *Server) {
		s.network = network
	}
}

func SetAddress(address string) ServerConfig {
	return func(s *Server) {
		s.address = address
	}
}

func SetTimeout(timeout time.Duration) ServerConfig {
	return func(s *Server) {
		s.timeout = timeout
	}
}
