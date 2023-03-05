package session_middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

type Service struct {
	log    abstract.Logger
	cfg    abstract.ConfigurationManager
	router *gin.RouterGroup

	sessionRouteRoot  *gin.RouterGroup
	sessionMiddleware gin.HandlerFunc
	config
	previousConfig string // to check when we need to reload
}

func (s *Service) Init(possibleDependencies *[]interface{}) {
	s.config.sessionOptions.HttpOnly = true
	s.config.sessionOptions.SameSite = http.SameSiteNoneMode
	s.config.sessionOptions.Secure = true
	s.config.sessionOptions.Path = "/"
	s.config.sessionOptions.MaxAge = secondsInDay * daysInWeek

	s.populateDependencies(possibleDependencies)
	s.loadConfig()
	go s.handleConfigChanges()
}

func (s *Service) Name() string {
	return "Session Middleware"
}
