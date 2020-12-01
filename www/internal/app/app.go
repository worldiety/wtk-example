package app

import (
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
	"github.com/golangee/forms-example/www/internal/tutorial/01-hello/01-helloworld"
	"github.com/golangee/forms-example/www/internal/tutorial/01-hello/02-hellohtml"
	"github.com/golangee/forms-example/www/internal/tutorial/01-hello/03-helloparam"
	"github.com/golangee/forms-example/www/internal/tutorial/01-hello/04-helloproperty"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
)

type Application struct {
	router *router.Router
	log    log.Logger
}

func NewApplication() *Application {
	a := &Application{
		router: router.NewRouter(),
		log:    log.NewLogger(ecs.Log("application")),
	}

	a.router.AddRoute("/", a.apply(a.index)).
		AddRoute(t01helloworld.Path, a.apply(t01helloworld.FromQuery)).
		AddRoute(t02hellohtml.Path, a.apply(t02hellohtml.FromQuery)).
		AddRoute(t03helloparam.Path, a.apply(t03helloparam.FromQuery)).
		AddRoute(t04helloproperty.Path, a.apply(t04helloproperty.FromQuery))

	a.router.
		SetUnhandledRouteAction(a.apply(func(query router.Query) Renderable {
			return Span(Text("unmatched route to " + query.Path()))
		}))

	return a
}

func (a *Application) apply(f func(query router.Query) Renderable) func(query router.Query) {
	return func(query router.Query) {
		RenderBody(a.page(query, f(query)))
	}
}

func (a *Application) index(q router.Query) Renderable {
	return P(Text("Welcome to the forms tutorial"))
}

func (a *Application) Run() {
	a.router.Start()
	select {}
}
