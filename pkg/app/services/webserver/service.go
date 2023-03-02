package webserver

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

type Service struct {
	log        abstract.Logger
	router     abstract.Router
	cfgManager abstract.ConfigurationManager
	server     *http.Server
	config     struct {
		port uint16
	}
}

func (s *Service) Init(allServices *[]interface{}) {
	s.populateDependencies(allServices)
	s.loadConfig()
	go s.loopUpdateConfig()
	go s.startWebServer()
}

func (s *Service) Name() string {
	return "Web Server"
}

func (s *Service) Msg(msg string) {
	s.log.Info().Msgf("[%s] %s", s.Name(), msg)
}

func (s *Service) Msgf(format string, v ...interface{}) {
	s.Msg(fmt.Sprintf(format, v...))
}

func (s *Service) loadConfig() {
	val := s.cfgManager.Get(s.Name(), "port")

	port, _ := strconv.ParseInt(val, 10, 64)

	if port <= 1 {
		port = 8080
	}

	s.config.port = uint16(port)
}

func (s *Service) loopUpdateConfig() {
	for {
		s.cfgManager.Set(s.Name(), "port", s.config.port)
		time.Sleep(time.Second)
	}
}

func (s *Service) startWebServer() {
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.port),
		Handler: s.router.RouteRoot(),
	}

	s.Msg("starting server ... ")
	if err := s.server.ListenAndServe(); err != nil {
	}
}
