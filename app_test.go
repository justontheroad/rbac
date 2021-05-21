package rbac

import "time"

func main() {
	app := New()
	time.AfterFunc(time.Second, func() {
		app.Stop()
	})
}
