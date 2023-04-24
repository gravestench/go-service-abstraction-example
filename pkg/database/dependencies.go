package database

import (
	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

func (s *Service) Dependencies() []interface{} {
	return []interface{}{s.cfg}
}

func (s *Service) ResolveDependencies(manager abstract.ServiceManager) {
	for _, other := range *manager.Services() {
		if cfg, ok := other.(abstract.ConfigurationManager); ok {
			s.cfg = cfg
		}
	}
}
