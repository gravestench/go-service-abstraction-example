package router

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/router/middleware"
)

func (s *Service) RouteRoot() *gin.Engine {
	return s.root
}

func (s *Service) createRouter() {
	s.Msg("creating root router instance")

	mode := gin.ReleaseMode
	if s.config.debug {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	s.root = gin.New()
	s.root.Use(middleware.Logger("gin", s.log))
}

func (s *Service) beginDynamicRouteBinding(allServices *[]interface{}) {
	for {
		s.bindNewRoutes(allServices)
		time.Sleep(time.Millisecond * 10)
	}
}

func (s *Service) bindNewRoutes(allServices *[]interface{}) {
	for _, candidate := range *allServices {
		svcToInit, ok := candidate.(abstract.RouteInitializer)
		if !ok {
			continue
		}

		if _, alreadyBound := s.boundServices[svcToInit.Name()]; alreadyBound {
			continue
		}

		s.Msgf("found new route initializer service %s", svcToInit.Name())

		if err := svcToInit.InitRoutes(s.root.Group(svcToInit.Slug())); err == nil {
			s.Msgf("bound routes for %s service", svcToInit.Name())
			s.boundServices[svcToInit.Name()] = nil // make 0-size entry
		}
	}
}
