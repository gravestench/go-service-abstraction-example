package webserver

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	}
}

func (s *Service) Init(allServices *[]interface{}) {
	s.populateDependencies(allServices)
	s.loadConfig()
	s.saveConfig()
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
		port = defaultPort
	}

	s.config.port = uint16(port)
}

func (s *Service) saveConfig() {
	if s.config.port == 0 {
		s.config.port = defaultPort
	}

	s.cfgManager.Set(s.Name(), "port", s.config.port)
}

func (s *Service) loopUpdateConfig() {
	for {
		s.handleChangedConfig()
		s.saveConfig()
		time.Sleep(time.Second)
	}
}

func (s *Service) startWebServer() {
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.port),
		Handler: s.router.RouteRoot(),
	}

	s.Msgf("starting server, listening on %d", s.config.port)
	if err := s.server.ListenAndServe(); err != nil {
	}
}

func (s *Service) stopWebServer() {
	s.Msg("stopping server ... ")

	_ = s.server.Close()
}

func (s *Service) handleChangedConfig() {
	strCurrentPort := fmt.Sprintf("%v", s.config.port)
	strConfigPort := s.cfgManager.Get(s.Name(), "port")

	if strCurrentPort == strConfigPort {
		return
	}

	port, _ := strconv.ParseInt(strConfigPort, 10, 64)
	s.config.port = uint16(port)

	s.Msg("port config has changed ...  ")
	s.stopWebServer()
	go s.startWebServer()
}
