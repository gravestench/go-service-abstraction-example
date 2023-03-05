package session_middleware

import (
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

func (s *Service) populateDependencies(others *[]interface{}) {
	for {
		if s.log != nil && s.cfg != nil && s.router != nil {
			break
		}

		for _, other := range *others {
			if logger, ok := other.(abstract.Logger); ok {
				s.log = logger
			}

			if cfg, ok := other.(abstract.ConfigurationManager); ok {
				s.cfg = cfg
			}

			if router, ok := other.(abstract.Router); ok {
				if root := router.RouteRoot(); root != nil {
					s.router = root.Group("")
				}
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
