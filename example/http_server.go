package main

import (
	// "time"

	"encoding/json"
	"net/http"

	"github.com/justontheroad/rbac/app"
	"github.com/justontheroad/rbac/handler"
	khttp "github.com/justontheroad/rbac/kernel/http"
	"github.com/justontheroad/rbac/middleware"
)

func main() {
	hs := khttp.NewServer(
		khttp.SetAddress(":8888"),
	)
	app := app.New(
		app.SetName("APP"),
		app.SetVersion("1.0.0"),
		app.SetServer(hs),
	)

	// use middleware
	hs.UseMiddleware(middleware.LoggingMiddleware)

	// handle func
	hs.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Subrouters
	sub := hs.Subrouters("/products")
	// "/products/"
	sub.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("producing"))
	})
	// "/products/{key}/"
	sub.HandleFunc("/{key}/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.RequestURI))
	})
	// "/products/{key}/details"
	sub.HandleFunc("/{key}/details", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("details"))
	})

	// Serving Single Page Applications
	spa := handler.SpaHandler{StaticPath: "public", IndexPath: "index.html"}
	hs.HandlePrefix("/", spa)

	// time.AfterFunc(time.Second, func() {
	// 	app.Stop()
	// })
	app.Run()
}
