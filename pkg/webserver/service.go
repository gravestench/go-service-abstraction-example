package webserver

import (
	"net/http"
	"strconv"

	"github.com/rs/zerolog"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

type Service struct {
	log        abstract.Logger
	router     abstract.Router
	cfgManager abstract.ConfigurationManager
	server     *http.Server
	lastConfig string
}

func (s *Service) Init(app abstract.ServiceManager) {
	s.ensureDefaults() // may be necessary to save the defaults
	go s.loopUpdateConfig()
	go s.startWebServer()
}

func (s *Service) SetLogger(l abstract.Logger) {
	s.log = l
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log.Logger()
}

func (s *Service) Name() string {
	return "Web Server"
}

func (s *Service) startWebServer() {
	tlsEnabled := s.Config().Get(keyTls) == "true"
	strPort := s.Config().Get(keyPort)
	port, _ := strconv.ParseInt(strPort, 10, 64)

	if port <= 1 {
		port = defaultPort
	}

	if tlsEnabled {
		s.log.Info().Msgf("starting HTTPS server, listening on %d", port)
		s.initTlsDebugServer()
		return
	}

	s.log.Info().Msgf("starting HTTP server, listening on %d", port)
	s.initHttpServer()
}

func (s *Service) stopWebServer() {
	if s.server == nil {
		return
	}

	s.log.Info().Msg("stopping server ... ")

	_ = s.server.Close()
}
