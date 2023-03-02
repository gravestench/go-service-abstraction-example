package webserver

import (
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

func (s *Service) populateDependencies(others *[]interface{}) {
	for {
		if s.log != nil && s.router != nil {
			break
		}

		for _, other := range *others {
			if logger, ok := other.(abstract.Logger); ok {
				s.log = logger
			}

			if router, ok := other.(abstract.Router); ok {
				if router.RouteRoot() != nil {
					s.router = router
				}
			}

			if cfg, ok := other.(abstract.ConfigurationManager); ok {
				s.cfgManager = cfg
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
