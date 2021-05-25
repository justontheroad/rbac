package middleware

import "context"

type Handler func(ctx context.Context, req interface{}) (interface{}, error)

type Middleware func(handler Handler) Handler

func Chain(outer Middleware, others ...Middleware) Middleware {
	return func(next Handler) Handler {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}
