package database

import (
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

func (s *Service) populateDependencies(others *[]interface{}) {
	for {
		if s.log != nil && s.cfg != nil {
			break
		}

		for _, other := range *others {
			if logger, ok := other.(abstract.Logger); ok {
				s.log = logger
			}

			if cfg, ok := other.(abstract.ConfigurationManager); ok {
				s.cfg = cfg.Group(s.Name())
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
