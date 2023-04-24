package backup_restore

import (
	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

func (s *Service) Dependencies() []interface{} {
	return []interface{}{s.cfg}
}

func (s *Service) ResolveDependencies(manager abstract.ServiceManager) {
	for _, other := range *manager.Services() {
		if config, ok := other.(abstract.ConfigurationManager); ok {
			s.cfg = config
		}
	}
}
