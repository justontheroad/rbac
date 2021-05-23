package app

import (
	"testing"
	"time"

	http "github.com/justontheroad/rbac/kernel/http"
)

func TestApp(t *testing.T) {
	hs := http.NewServer()
	app := New(
		SetName("APP"),
		SetVersion("1.0.0"),
		SetServer(hs),
	)
	time.AfterFunc(time.Second, func() {
		app.Stop()
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
