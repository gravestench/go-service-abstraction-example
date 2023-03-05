package static_assets

import (
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

func (s *Service) populateDependencies(others *[]interface{}) (keepLooking bool) {
	if s.log != nil && s.router != nil {
		return false
	}

	for _, other := range *others {
		if logger, ok := other.(abstract.Logger); ok {
			s.log = logger
		}

		if router, ok := other.(abstract.Router); ok {
			if root := router.RouteRoot(); root != nil {
				s.router = router
			}
		}
	}

	return true
}
