package middleware

import (
	"log"
	"net/http"
	"net/url"
)

type HTTPLogger struct {
	level     uint16
	path      string
	method    string
	args      url.Values
	component string
	query     string
}

func (logger *HTTPLogger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := &HTTPLogger{
			level:     200,
			path:      r.URL.Path,
			method:    r.Method,
			args:      r.Form,
			component: "HTTP",
			query:     r.URL.RawQuery,
		}
		log.Println(l)
		next.ServeHTTP(w, r)
	})
}
