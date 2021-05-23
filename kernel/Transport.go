package kernel

import "context"

type Server interface {
	Start() error
	Stop() error
}

type Kind string

type Transport struct {
	Kind Kind
}

const (
	Http Kind = "HTTP"
)

type transportKey struct{}

func NewContext(cxt context.Context, tr Transport) context.Context {
	return context.WithValue(cxt, transportKey{}, tr)
}

func fromContext(cxt context.Context) (tr Transport, ok bool) {
	tr, ok = cxt.Value(transportKey{}).(Transport)
	return
}
