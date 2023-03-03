package webserver

import (
	"fmt"
	"strconv"
	"time"
)

func (s *Service) loadConfig() {
	val := s.cfgManager.Get(s.Name(), "port")

	port, _ := strconv.ParseInt(val, 10, 64)

	if port <= 1 {
		port = defaultPort
	}

	s.config.port = uint16(port)
	s.config.tls = s.cfgManager.Get(s.Name(), "tls") == "true"
}

func (s *Service) saveConfig() {
	if s.config.port == 0 {
		s.config.port = defaultPort
	}

	s.cfgManager.Set(s.Name(), "port", s.config.port)

	tls := "false"
	if s.config.tls {
		tls = "true"
	}
	s.cfgManager.Set(s.Name(), "tls", tls)
}

func (s *Service) loopUpdateConfig() {
	for {
		s.handleChangedConfig()
		s.saveConfig()
		time.Sleep(time.Second)
	}
}

func (s *Service) handleChangedConfig() {
	var hasChanged bool

	defer func() {
		if !hasChanged {
			return
		}

		s.stopWebServer()
		go s.startWebServer()
	}()

	strCurrentPort := fmt.Sprintf("%v", s.config.port)
	strConfigPort := s.cfgManager.Get(s.Name(), "port")

	if strCurrentPort != strConfigPort {
		s.Msg("port config has changed ...  ")

		hasChanged = true
		port, _ := strconv.ParseInt(strConfigPort, 10, 64)
		s.config.port = uint16(port)
	}

	strCurrentTls := "false"
	if s.config.tls {
		strCurrentTls = "true"
	}
	strConfigTls := s.cfgManager.Get(s.Name(), "tls")

	if strCurrentTls != strConfigTls {
		s.Msg("TLS config has changed ...  ")

		hasChanged = true
		s.config.tls = strConfigTls == "true"
	}
}
