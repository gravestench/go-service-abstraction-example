package webserver

import (
	"fmt"
	"net/http"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

const (
	defaultPort = 8080
)

type Service struct {
	log        abstract.Logger
	router     abstract.Router
	cfgManager abstract.ConfigurationManager
	server     *http.Server
	config     struct {
		port uint16
		tls  bool
	}
}

func (s *Service) Init(allServices *[]interface{}) {
	s.populateDependencies(allServices)
	s.loadConfig() // may populate defaults
	s.saveConfig() // may be necessary to save the defaults
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

func (s *Service) startWebServer() {
	if s.config.tls {
		s.Msgf("starting HTTPS server, listening on %d", s.config.port)
		s.initTlsDebugServer()
		return
	}

	s.Msgf("starting HTTP server, listening on %d", s.config.port)
	s.initHttpServer()
}

func (s *Service) stopWebServer() {
	s.Msg("stopping server ... ")

	_ = s.server.Close()
}
