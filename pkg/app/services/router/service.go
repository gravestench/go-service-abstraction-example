package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

type Service struct {
	root          *gin.Engine
	log           abstract.Logger
	cfgManager    abstract.ConfigurationManager
	boundServices map[string]*struct{} // holds 0-size entries
	config        struct {
		debug bool
	}
}

func (s *Service) Init(allServices *[]interface{}) {
	s.boundServices = make(map[string]*struct{})
	s.populateDependencies(allServices)
	s.loadConfig()
	s.createRouter()
	go s.beginDynamicRouteBinding(allServices)
	go s.handleConfigChanges()
}

func (s *Service) Name() string {
	return "Route Manager"
}

func (s *Service) Msg(msg string) {
	s.log.Info().Msgf("[%s] %s", s.Name(), msg)
}

func (s *Service) Msgf(format string, v ...interface{}) {
	s.Msg(fmt.Sprintf(format, v...))
}
