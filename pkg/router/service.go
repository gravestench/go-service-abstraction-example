package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/router/middleware/cors"
	"github.com/gravestench/go-service-abstraction-example/pkg/router/middleware/session"
	"github.com/gravestench/go-service-abstraction-example/pkg/router/middleware/static_assets"
)

type Service struct {
	app        abstract.ServiceManager
	log        abstract.Logger
	cfgManager abstract.ConfigurationManager

	cors    *cors.Middleware
	session *session.Middleware
	static  *static_assets.Middleware

	root      *gin.Engine
	protected *gin.RouterGroup

	boundServices map[string]*struct{} // holds 0-size entries

	config struct {
		debug bool
	}

	flags struct {
		parsed          bool
		slug            string
		disableSessions bool
		logLevel        int
	}
}

func (s *Service) SetLogger(l abstract.Logger) {
	s.log = l
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log.Logger()
}

func (s *Service) Init(app abstract.ServiceManager) {
	s.app = app
	s.loadConfig()
	s.dwellUntilFlagsParsed()
	go s.beginDynamicRouteBinding()
	go s.handleConfigChanges()
}

func (s *Service) Name() string {
	return "Route Manager"
}
