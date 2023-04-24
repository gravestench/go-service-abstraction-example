package router

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"k8s.io/utils/strings/slices"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/router/middleware/cors"
	"github.com/gravestench/go-service-abstraction-example/pkg/router/middleware/logger"
	"github.com/gravestench/go-service-abstraction-example/pkg/router/middleware/session"
	"github.com/gravestench/go-service-abstraction-example/pkg/router/middleware/static_assets"
)

var _ abstract.Router = &Service{}

func (s *Service) RouteRoot() *gin.Engine {
	return s.root
}

func (s *Service) RouteProtected() *gin.RouterGroup {
	return s.protected
}

func (s *Service) initRouter() {
	for {
		if s.flags.parsed {
			break
		}

		time.Sleep(time.Second)
	}

	if s.boundServices == nil {
		s.log.Info().Msg("initializing routes")
	} else {
		s.log.Info().Msg("re-initializing routes")
	}

	// forget any already-bound services
	s.boundServices = make(map[string]*struct{})

	mode := gin.ReleaseMode
	if s.config.debug || s.flags.logLevel <= int(zerolog.DebugLevel) {
		mode = gin.DebugMode
	}

	// set up the gin router
	gin.SetMode(mode)
	s.root = gin.New()
	s.root.Use(logger.Logger("gin", s.log))

	// set up the middleware
	s.session = session.New(s.log, s.cfgManager)
	s.cors = cors.New(s.log, s.cfgManager)

	// use the CORS middleware
	s.root.Use(s.cors.GetCorsHandler())

	// set up a protected route group
	s.protected = s.root.Group("api")

	if !s.flags.disableSessions {
		// use the session middleware for the protected route group
		for {
			if sm := s.session.SessionHandler(); sm != nil {
				s.protected.Use(sm)
				break
			}
		}

		s.protected.Use(s.session.Checkauth)
		s.protected.Use(s.session.CheckRole)
	}

	s.static = static_assets.New(s.log, s)
}

func (s *Service) beginDynamicRouteBinding() {
	for {
		if s.shouldInit() {
			s.initRouter()
		}

		s.bindNewRoutes()
		time.Sleep(time.Millisecond * 10)
	}
}

func (s *Service) shouldInit() bool {
	if s.boundServices == nil {
		return true // base case, happens at app start
	}

	// in the event that a service is removed by the
	// service manager for whatever reason, we need to check
	// if that was something that had routes. if it was, we need
	// to re-init the router (we can't actually remove routes in gin)

	// we will check if any of the services that have routes are no longer
	// in the list of the service managers services
	allExistingServiceNames := make([]string, 0)
	for _, candidate := range *s.app.Services() {
		if svc, ok := candidate.(abstract.RouteInitializer); ok {
			allExistingServiceNames = append(allExistingServiceNames, svc.Name())
			continue
		}

		if svc, ok := candidate.(abstract.ProtectedRouteInitializer); ok {
			allExistingServiceNames = append(allExistingServiceNames, svc.Name())
			continue
		}
	}

	// iterate over each bound service, check if its name
	// exists as a substring inside of our lookup string
	for key, _ := range s.boundServices {
		if !slices.Contains(allExistingServiceNames, key) {
			return true
		}
	}

	return false
}

func (s *Service) bindNewRoutes() {
	for _, candidate := range *s.app.Services() {
		svcToInit, ok := candidate.(abstract.Service)
		if !ok {
			continue
		}

		if _, alreadyBound := s.boundServices[svcToInit.Name()]; alreadyBound {
			continue
		}

		group := ""
		if svc, ok := candidate.(abstract.HasRouteSlug); ok {
			group = svc.Slug()
		}

		// handle unprotected route init
		if r, ok := candidate.(abstract.RouteInitializer); ok {
			s.log.Info().Msgf("binding routes for the '%s' service", svcToInit.Name())
			if s.flags.slug != "" {
				group = fmt.Sprintf("%s/%s", s.flags.slug, group)
			}

			if err := r.InitRoutes(s.root.Group(group)); err == nil {
				s.boundServices[svcToInit.Name()] = nil // make 0-size entry
			}

			continue
		}

		// handle protected route init
		if r, ok := candidate.(abstract.ProtectedRouteInitializer); ok {
			s.log.Info().Msgf("binding protected routes for the '%s' service", svcToInit.Name())
			if s.flags.slug != "" {
				group = fmt.Sprintf("%s/%s", s.flags.slug, group)
			}

			if err := r.InitProtectedRoutes(s.protected.Group(group)); err == nil {
				s.boundServices[svcToInit.Name()] = nil // make 0-size entry
			}

			continue
		}
	}
}
