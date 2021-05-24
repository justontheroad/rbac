package example

import (
	"time"

	"github.com/justontheroad/rbac/app"
	"github.com/justontheroad/rbac/kernel/http"
)

func main() {
	hs := http.NewServer(
		http.SetAddress(":8888"),
	)
	app := app.New(
		app.SetName("APP"),
		app.SetVersion("1.0.0"),
		app.SetServer(hs),
	)
	time.AfterFunc(time.Second, func() {
		app.Stop()
	})
	app.Run()
}
