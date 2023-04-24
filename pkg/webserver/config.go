package webserver

import (
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

const (
	keyPort     = "port"
	defaultPort = 8080

	keyTls     = "tls"
	defaultTls = false
)

func (s *Service) Config() abstract.ConfigGroup {
	return s.cfgManager.GetConfigGroup(s.Name())
}

func (s *Service) ensureDefaults() {
	s.Config().Default(keyPort, defaultPort)
	s.Config().Default(keyTls, defaultTls)
}

func (s *Service) loopUpdateConfig() {
	for {
		s.ensureDefaults()
		s.handleChangedConfig()
		time.Sleep(time.Second)
	}
}

func (s *Service) handleChangedConfig() {
	currentConfig := s.Config().String()

	if s.lastConfig == "" {
		s.lastConfig = currentConfig
	}

	if s.lastConfig == currentConfig {
		return
	}

	s.stopWebServer()
	go s.startWebServer()
}
