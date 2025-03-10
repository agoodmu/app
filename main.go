package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type hello struct {
	app.Compo

	name string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello, "),
			app.If(h.name != "", func() app.UI {
				return app.Text(h.name)
			}).Else(func() app.UI {
				return app.Text("World!")
			}),
		),
		app.P().Body(
			app.Input().
				Type("text").
				Value(h.name).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(h.ValueTo(&h.name)),
		),
	)
}

func main() {
	// Go-app component routing (client-side):
	app.Route("/", func() app.Composer { return &hello{name: "go"} })
	app.Route("/hello", func() app.Composer { return &hello{name: "hello"} })
	app.RunWhenOnBrowser()

	// Standard HTTP routing (server-side):
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
