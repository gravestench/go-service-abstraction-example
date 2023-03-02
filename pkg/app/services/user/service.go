package user

import (
	"fmt"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

type Service struct {
	log abstract.Logger
	cfg abstract.ConfigurationManager
}

func (s *Service) Name() string {
	return "User Management"
}

func (s *Service) Init(dependencies *[]interface{}) {
	s.populateServices(dependencies)
	go s.start()
}

func (s *Service) start() {
	s.Msg("starting")
}

func (s *Service) Msg(msg string) {
	s.log.Info().Msgf("[%s] %s", s.Name(), msg)
}

func (s *Service) Msgf(format string, v ...interface{}) {
	s.log.Info().Msg(fmt.Sprintf(format, v...))
}
