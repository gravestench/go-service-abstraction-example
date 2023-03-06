package flags_manager

import (
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

func (s *Service) populateDependencies(others *[]interface{}) {
	for {
		if s.log != nil {
			break
		}

		for _, other := range *others {
			if logger, ok := other.(abstract.Logger); ok {
				s.log = logger
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
