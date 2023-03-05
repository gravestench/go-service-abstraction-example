package static_assets

import (
	"time"

	"github.com/gin-contrib/gzip"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

type Service struct {
	log    abstract.Logger
	router abstract.Router
}

func (s *Service) Name() string {
	return "Static Assets"
}

func (s *Service) Init(possibleDependencies *[]interface{}) {
	const pollingInterval = time.Second

	for s.populateDependencies(possibleDependencies) {
		time.Sleep(pollingInterval)
	}

	s.router.RouteRoot().Use(gzip.Gzip(gzip.DefaultCompression))
	s.router.RouteRoot().NoRoute(s.staticWebUIHandler)
}
