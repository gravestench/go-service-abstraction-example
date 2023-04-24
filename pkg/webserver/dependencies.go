package webserver

import (
	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

func (s *Service) Dependencies() []interface{} {
	return []interface{}{s.router, s.cfgManager}
}

func (s *Service) ResolveDependencies(manager abstract.ServiceManager) {
	for _, other := range *manager.Services() {
		if router, ok := other.(abstract.Router); ok {
			if router.RouteRoot() != nil {
				s.router = router
			}
		}

		if cfg, ok := other.(abstract.ConfigurationManager); ok {
			s.cfgManager = cfg
		}
	}
}
