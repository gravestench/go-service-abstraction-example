package config

import (
	"fmt"
	"sync"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

// each service will maintain its own config
type perServiceConfig = map[string]map[string]string

type Service struct {
	store perServiceConfig
	log   abstract.Logger
	mux   sync.Mutex
}

func (s *Service) Init(allServices *[]interface{}) {
	s.populateDependencies(allServices)

	s.store = make(perServiceConfig)

	s.LoadConfig()
	go s.loopSaveConfig()
}

func (s *Service) Name() string {
	return "Configuration Manager"
}

func (s *Service) Msg(msg string) {
	s.log.Info().Msgf("[%s] %s", s.Name(), msg)
}

func (s *Service) Msgf(format string, v ...interface{}) {
	s.log.Info().Msg(fmt.Sprintf(format, v...))
}
