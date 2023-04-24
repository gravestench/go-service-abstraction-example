package static_assets

import (
	"github.com/gin-contrib/gzip"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

func New(log abstract.Logger, r abstract.Router) *Middleware {
	m := &Middleware{
		log:    log,
		router: r,
	}

	m.init()

	return m
}

type Middleware struct {
	log    abstract.Logger
	router abstract.Router
}

func (m *Middleware) Name() string {
	return "Static Assets Middleware"
}

func (m *Middleware) init() {
	m.log.Info().Msgf("[%s] setting up routes", m.Name())
	m.router.RouteRoot().Use(gzip.Gzip(gzip.DefaultCompression))
	m.router.RouteRoot().NoRoute(m.staticWebUIHandler)
}
